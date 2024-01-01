package test

import (
	"fintechpractices/configs"
	"fintechpractices/global"
	Init "fintechpractices/init"
	"fintechpractices/internal/controller"
	"fintechpractices/internal/task"
	"fintechpractices/internal/task/types"
	"fmt"
	"reflect"
	"testing"
	"time"
)

type TestTaskArgs struct {
	CmdName string
	Args1   string
	Args2   string
}

func (t TestTaskArgs) Map() map[string]interface{} {
	return t.struct2map(t)
}

func (t TestTaskArgs) TaskName() string {
	return "testTask"
}

func (t TestTaskArgs) struct2map(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	reflectValue := reflect.ValueOf(data)
	if reflectValue.Kind() != reflect.Struct {
		return nil
	}
	reflectType := reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fieldName := reflectType.Field(i).Name
		fieldValue := reflectValue.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result
}

func InitAndStart() {
	Init.Initialization()

	task.TaskToEntity[TestTaskArgs{}.TaskName()] = configs.ModelEntity{
		BeforeStart: `C:\GoProject\fintech_practices\test\testTask.exe --method "before_cmd_execute"`,
		ExecCmd:     "{{.CmdName}} {{.Args1}} {{.Args2}}",
		AfterEnd:    `C:\GoProject\fintech_practices\test\testTask.exe --method "after_cmd_execute"`,
	}

	go global.TaskMgr.Start()
}

func TestTaskMgr(t *testing.T) {
	InitAndStart()

	err := global.TaskMgr.RegisterTask("test_Task_1", TestTaskArgs{
		CmdName: `C:\GoProject\fintech_practices\test\testTask.exe`,
		Args1:   "--method",
		Args2:   `test_Task_1`,
		// Args2: "",
	})
	if err != nil {
		t.Errorf("task <test_Task_1> register error: %s", err.Error())
		return
	}
	fmt.Println("=====task register success=====")
	for {
		data, ok := global.TaskMgr.QueryTask("test_Task_1")
		if !ok {
			t.Errorf("can not find task")
			return
		}
		info := data.(task.TaskInfo)
		if info.Status == types.Runnable || info.Status == types.Running {
			time.Sleep(time.Second * 3)
			fmt.Printf("task <test_Task_1> is %d\n", info.Status)
			continue
		}

		if info.Status == types.Failed {
			t.Errorf(info.Msg)
			return
		}
		fmt.Println(info)
		return
	}
}

func TestVitsModel(t *testing.T) {
	InitAndStart()

	err := global.TaskMgr.RegisterTask("test_audio_task", types.AudioArgs{
		TextInput: "尝试生成一个测试音频任务，并进行执行",
		ToneInput: "update_qihan",
		OutputDir: global.RootDirMap[controller.FtypeAudio.String()],
		FileName:  "test_audio_task.wav",
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	for {
		data, ok := global.TaskMgr.QueryTask("test_audio_task")
		if !ok {
			t.Errorf("can not find task")
			return
		}
		info := data.(task.TaskInfo)
		if info.Status == types.Runnable || info.Status == types.Running {
			time.Sleep(time.Second * 3)
			fmt.Printf("task <test_audio_task> is %d\n", info.Status)
			continue
		}

		if info.Status == types.Failed {
			t.Errorf(info.Msg)
			return
		}
		fmt.Println(info)
		return
	}
}
