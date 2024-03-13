package config

type Logger struct {
	Out   bool   `yaml:"out"`
	Level string `yaml:"loglevel"`
	Path  string `yaml:"logpath"`
	Name  string `yaml:"logname"`
	Size  int64  `yaml:"logsize"`
}
