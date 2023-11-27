package configs

type LogConfig struct {
	FileName   string `mapstructure:"filename" yaml:"filename"`
	Level      string `mapstructure:"level" yaml:"level"`
	Dir        string `mapstructure:"dir" yaml:"dir"`
	LineNo     bool   `mapstructure:"line_no" yaml:"line_no"`
	Format     string `mapstructure:"format" yaml:"format"`
	MaxBackups int    `mapstructure:"max_backups" yaml:"max_backups"`
	MaxSize    int    `mapstructure:"max_size" yaml:"max_size"`
	MaxAge     int    `mapstructure:"max_age" yaml:"max_age"`
}
