package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

func init() {
	workdir, _ := os.Getwd()
	Conf := &Con
	viper.SetConfigFile(workdir + "/client.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("setting init error %v", err)
	}

	if err := viper.Unmarshal(Conf); err != nil {
		log.Panicf("setting init error %v", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("文件改变啦")
	})
}
