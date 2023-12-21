package configs

type ServerConfig struct {
	Host              string `mapstructure:"host" yaml:"host"`
	Port              int    `mapstructure:"port" yaml:"port"`
	EngineMode        string `mapstructure:"engine_mode" yaml:"engine_mode"`
	DpRootDir         string `mapstructure:"dp_root_dir" yaml:"dp_root_dir"`
	ResourceRootDir   string `mapstructure:"resource_root_dir" yaml:"resource_root_dir"`
	CoverImageRootDir string `mapstructure:"cover_image_root_dir" yaml:"cover_image_root_dir"`
	AudioRootDir      string `mapstructure:"audio_root_dir" yaml:"audio_root_dir"`
}
