package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

func initConfigs() {
	var configFilePath string
	viper.SetConfigFile("config")
	if configFilePath != "" {
		stat, err := os.Stat(configFilePath)
		if err != nil {
			fmt.Println("Error reading config file ,%s", err)
		}
		if stat.IsDir() {
			viper.AddConfigPath(configFilePath)
		} else {
			viper.AddConfigPath(configFilePath)
		}
	}
	viper.AddConfigPath(".")
	viper.AddConfigPath(os.Getenv("/etc/appname/")) // path to look for the config file in
	viper.AddConfigPath(os.Getenv("$HOME/.appname"))
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file ,%s", err)
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s\n", e.Name)
	})
	viper.WatchConfig()
}

type ViperConfig struct {
}

func (v ViperConfig) build() {
	initConfigs()
}

func (v ViperConfig) IsSet(key string) bool {
	return viper.IsSet(key)
}

func (v ViperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v ViperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v ViperConfig) GetFloat(key string) float64 {
	return viper.GetFloat64(key)
}

func (v ViperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (v ViperConfig) GetIntSlice(key string) []int {
	return viper.GetIntSlice(key)
}

func (v ViperConfig) GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func (v ViperConfig) GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

func (v ViperConfig) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}
