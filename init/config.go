package init

import (
	"flag"

	"fintechpractices/global"

	"github.com/spf13/viper"
)

func InitConfig() error {
	cfgPath := flag.String("config", "", "config file")
	flag.Parse()
	if *cfgPath != "" {
		viper.SetConfigFile(*cfgPath)
		viper.SetConfigType("yaml")
	} else {
		viper.SetConfigFile("../conf/conf.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(global.AppCfg); err != nil {
		return err
	}

	return nil
}
