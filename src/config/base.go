package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

func init() {
	absPath, err := filepath.Abs("./")
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}
	viper.SetConfigName("dev.config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(absPath)
	err = viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
		panic(err)
	}
}
