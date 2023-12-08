package main

import (
	"fintechpractices/global"
	Init "fintechpractices/init"
	"fmt"
)

//	@title			cc fintech practices API
//	@version		1.0
//	@description	榕树平台API
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	rliu
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost:8080

func main() {

	Init.Initialization()
	Run()
}

func Run() {
	if err := global.Engine.Run(fmt.Sprintf("%s:%d", global.AppCfg.ServerCfg.Host, global.AppCfg.ServerCfg.Port)); err != nil {
		panic(err.Error())
	}
}
