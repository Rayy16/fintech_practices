package test

import (
	"fintechpractices/global"
	Init "fintechpractices/init"
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	if err := Init.InitConfig(); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(global.AppCfg)

	// cfgPath := "../conf/conf.yaml"
	// viper.SetConfigFile(cfgPath)

	// if err := viper.ReadInConfig(); err != nil {
	// 	t.Error(err.Error())
	// }
	// appCfg := configs.AppConfig{}
	// if err := viper.Unmarshal(&appCfg); err != nil {
	// 	t.Error(appCfg)
	// }

	// fmt.Printf("%#v\n", appCfg)
}

func TestInitZaplog(t *testing.T) {
	if err := Init.InitConfig(); err != nil {
		t.Error(err.Error())
	}

	if log := Init.InitZapLog(); log == nil {
		t.Error("init zap log failed")
	} else {
		log := log.Sugar()
		log.Debug("hello world")
		log.Info("hello world!")
		log.Error("hello world!")
		log.Warn("hello world!")
		// log.Panic("hello world!")
		// log.Fatal("hello world!")
	}

}

func TestInitMySQLGorm(t *testing.T) {
	if err := Init.InitConfig(); err != nil {
		t.Error(err.Error())
	}

	db := Init.InitMySQLGorm()
	if db == nil {
		t.Error("init mysql gorm db failed")
	}

	if db := db.Exec("show databases;"); db.Error != nil {
		t.Error(db.Error.Error())
	} else {
		fmt.Println(db.RowsAffected)
	}
}
