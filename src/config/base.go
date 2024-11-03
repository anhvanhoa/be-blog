package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func init() {
	absPath, err := filepath.Abs("./")
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}
	mode := os.Getenv("ENV_MODE")
	if mode == "dev" {
		viper.SetConfigName("dev.config")
	} else if mode == "production" {
		viper.SetConfigName("production.config")
	} else {
		panic("ENV_MODE is invalid")
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath(absPath)
	err = viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
		panic(err)
	}
}
