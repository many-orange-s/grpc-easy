package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

func Init() {
	workdir, _ := os.Getwd()
	Conf := &Con
	viper.SetConfigFile(workdir + "/service.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("setting init error %v", err)
	}
	if err = viper.Unmarshal(Conf); err != nil {
		log.Panicf("setting init error %v", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Print("文件修改了")
	})
}
