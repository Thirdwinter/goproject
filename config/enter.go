// !config包存放配置信息
package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
	Redis  Redis  `yame:"redis"`
}
