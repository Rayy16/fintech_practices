package main

import (
	"fintechpractices/global"
	Init "fintechpractices/init"
	"fmt"
)

func main() {

	Init.Initialization()
	Run()
}

func Run() {
	if err := global.Engine.Run(fmt.Sprintf("%s:%d", global.AppCfg.ServerCfg.Host, global.AppCfg.ServerCfg.Port)); err != nil {
		panic(err.Error())
	}
}
