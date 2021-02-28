package tools

import (
	"github.com/spf13/viper"
	"log"
)

func GetEnv(key, defaultValue string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Keyname : %v, not found, default key value : %v, has been loaded", key, defaultValue)
		return defaultValue
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}