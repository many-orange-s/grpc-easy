package config

type Config struct {
	Port    string `json:"port" mapstructure:"port"`
	CarFile string `json:"carFile" mapstructure:"carFile"`
	KeyFile string `json:" keyFile" mapstructure:"keyFile"`
	CaFile  string `json:"caFile" mapstructure:"caFile"`
}

var Con Config
