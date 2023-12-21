package init

import (
	"fintechpractices/configs"
	"fintechpractices/global"
	"fintechpractices/internal/task"
	"fintechpractices/internal/task/types"
)

func InitTaskMgr() types.TaskManagerIntf {
	initTaskCentra()
	return task.NewTaskMgr(1000)
}

func initTaskCentra() {
	task.TaskToEntity = make(map[string]configs.ModelEntity)
	task.TaskToEntity[types.AudioArgs{}.TaskName()] = global.AppCfg.ModelCfg.AudioModel
	task.TaskToEntity[types.DpArgs{}.TaskName()] = global.AppCfg.ModelCfg.DpModel
}
