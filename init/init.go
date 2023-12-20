package init

import (
	"fintechpractices/global"
)

func Initialization() {
	err := InitConfig()
	if err != nil {
		panic(err.Error())
	}

	global.Log = InitZapLog()
	global.DB = InitMySQLGorm()
	global.Engine = InitServer()
	global.TaskMgr = InitTaskMgr()
}
