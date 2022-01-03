package config

type Config struct {
	Address  string `json:"address" mapstructure:"address"`
	Hostname string `json:"hostname" mapstructure:"hostname"`
	CrtFile  string `json:"crtFile" mapstructure:"crtFile"`
	KeyFile  string `json:"kayFile" mapstructure:"keyFile"`
	CaFile   string `json:"caFile" mapstructure:"caFile"`
	Secrete  string `json:"secrete" mapstructure:"secrete"`
}

var Con Config
