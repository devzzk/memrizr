package config

type Config struct {
	Mysql  Mysql `yaml:"mysql"`
	Logger Mysql `yaml:"logger"`
	System Mysql `yaml:"system"`
}
