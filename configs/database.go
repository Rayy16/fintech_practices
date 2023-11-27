package configs

type DataBaseConfig struct {
	Host          string `mapstructure:"host" yaml:"host"`
	Port          int    `mapstructure:"port" yaml:"port"`
	User          string `mapstructure:"user" yaml:"user"`
	Password      string `mapstructure:"password" yaml:"password"`
	Database      string `mapstructure:"database" yaml:"database"`
	LogToFile     bool   `mapstructure:"log_to_file" yaml:"log_to_file"`
	FileName      string `mapstructure:"filename" yaml:"filename"`
	LogMode       string `mapstructure:"log_mode" yaml:"log_mode"`
	SlowThreshold int    `mapstructure:"slow_threshold" yaml:"slow_threshold"`
	MaxIdleConns  int    `mapstructure:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns  int    `mapstructure:"max_open_conns" yaml:"max_open_conns"`
}
