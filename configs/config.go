package configs

type AppConfig struct {
	DbCfg     DataBaseConfig `mapstructure:"database" yaml:"database"`
	LogCfg    LogConfig      `mapstructure:"log" yaml:"log"`
	ServerCfg ServerConfig   `mapstructure:"server" yaml:"server"`
	ModelCfg  ModelConfig    `mapstructure:"model" yaml:"model"`
}
