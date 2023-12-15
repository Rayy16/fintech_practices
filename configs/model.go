package configs

type ModelConfig struct {
	DpModel    ModelEntity `mapstructure:"dp" yaml:"dp"`
	AudioModel ModelEntity `mapstructure:"audio" yaml:"audio"`
}

type ModelEntity struct {
	ExecCmd     string `mapstructure:"exec_cmd" yaml:"exec_cmd"`
	BeforeStart string `mapstructure:"before_start" yaml:"before_start"`
	AfterEnd    string `mapstructure:"after_end" yaml:"after_end"`
}
