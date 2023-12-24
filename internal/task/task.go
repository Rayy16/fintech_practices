package task

import (
	"bytes"
	"fintechpractices/configs"
	"fintechpractices/global"
	"fintechpractices/internal/dao"
	"fintechpractices/internal/task/types"
	"fintechpractices/tools"
	"fmt"
	"strings"
	"sync"
	"text/template"
	"time"
)

var TaskToEntity map[string]configs.ModelEntity

type TaskInfo struct {
	Status int
	Msg    string
}

type ExecCommand struct {
	BeforeCmd string
	Cmd       string
	AfterCmd  string
}

type TaskExcutor struct {
	TaskId string
	Cmds   []*ExecCommand
	Err    error
}

func (e *TaskExcutor) Execute() {
	var std [2]string
	for _, execCmd := range e.Cmds {
		if e.Error() != nil {
			return
		}
		e.executeStr(execCmd.BeforeCmd)
		std, e.Err = e.executeStr(execCmd.Cmd)
		e.executeStr(execCmd.AfterCmd)
	}
	if e.Error() != nil {
		return
	}
	e.Err = e.checkResult(std[0])
}

func (e *TaskExcutor) executeStr(cmd string) (std [2]string, err error) {
	if cmd == "" {
		return [2]string{}, nil
	}
	log := global.Log.Sugar()
	log.Infof("exec cmd: %s", cmd)
	var name string
	var args []string
	name, args = tools.SplitCmd(cmd)
	return tools.RunCmd(name, args)

}

func (e *TaskExcutor) Error() error {
	return e.Err
}

func (e *TaskExcutor) GetTaskId() string {
	return e.TaskId
}

func (e *TaskExcutor) checkResult(res string) error {
	if !strings.Contains(res, "The generated video is named") ||
		!strings.Contains(res, e.GetTaskId()) {
		return fmt.Errorf("task <%s> not match expection", e.GetTaskId())
	}
	return nil
}

type TaskManager struct {
	taskCh     chan TaskExcutor
	taskStatus sync.Map
}

func NewTaskMgr(cocurrent int) *TaskManager {
	return &TaskManager{
		taskCh:     make(chan TaskExcutor, cocurrent),
		taskStatus: sync.Map{},
	}
}

func (m *TaskManager) RegisterTask(taskID string, args ...types.TaskArgs) error {
	log := global.Log.Sugar()

	if status, ok := m.taskStatus.Load(taskID); ok {
		info := status.(TaskInfo)
		dpStatus := info.Status
		if dpStatus == types.Running || dpStatus == types.Runnable {
			log.Info("task already register, task status: %d", dpStatus)
			return nil
		}
	}

	execCmds, err := m.convertToCmd(args...)
	if err != nil {
		log.Errorf("convertToCmd error: %s", err.Error())
		return err
	}
	excutor := TaskExcutor{
		taskID,
		execCmds,
		nil,
	}
	m.taskStatus.Store(
		taskID,
		TaskInfo{dao.StatusCreatable.Int(), fmt.Sprintf("task <%s> register success", taskID)},
	)
	go func() {
		m.taskCh <- excutor
	}()
	return nil
}

func (m *TaskManager) UpdateTask(taskID string, data interface{}) {
	log := global.Log.Sugar()
	info := data.(TaskInfo)
	m.taskStatus.Store(taskID, info)
	log.Infof("update task <%s> to <%+v>", taskID, data)
}

func (m *TaskManager) convertToCmd(args ...types.TaskArgs) ([]*ExecCommand, error) {
	var cmdBuf bytes.Buffer
	log := global.Log.Sugar()
	cmds := make([]*ExecCommand, 0, len(args))
	for _, a := range args {
		log.Infof("convert args to command, args: %v", a)
		entity := TaskToEntity[a.TaskName()]
		tmpl, err := template.New("modelCommand").Parse(entity.ExecCmd)
		if err != nil {
			log.Errorf("template.New().Parse error: %s", err.Error())
			return nil, err
		}
		err = tmpl.Execute(&cmdBuf, a.Map())
		if err != nil {
			log.Errorf("tmpl.Execute error: %s", err.Error())
			return nil, err
		}
		execCmd := &ExecCommand{
			BeforeCmd: entity.BeforeStart,
			Cmd:       cmdBuf.String(),
			AfterCmd:  entity.AfterEnd,
		}
		cmds = append(cmds, execCmd)
	}
	return cmds, nil
}

func (m *TaskManager) rangeTasks() types.TaskExcutorIntf {
	executor, ok := <-m.taskCh
	if !ok {
		return nil
	}
	return &executor
}

func (m *TaskManager) Start() {
	log := global.Log.Sugar()
	for {
		executor := m.rangeTasks()
		if executor == nil {
			log.Infof("task manager is closed at %v", time.Now())
			return
		}
		go func() {
			m.UpdateTask(executor.GetTaskId(),
				TaskInfo{types.Running, fmt.Sprintf("execute task <%s> at %v", executor.GetTaskId(), time.Now())},
			)
			executor.Execute()
			if err := executor.Error(); err != nil {
				m.UpdateTask(executor.GetTaskId(),
					TaskInfo{types.Failed, err.Error()},
				)
			} else {
				m.UpdateTask(executor.GetTaskId(),
					TaskInfo{types.Success, fmt.Sprintf("task <%s> execute success at %v", executor.GetTaskId(), time.Now())},
				)
			}
		}()
	}
}

func (m *TaskManager) QueryTask(taskId string) (interface{}, bool) {
	return m.taskStatus.Load(taskId)
}
