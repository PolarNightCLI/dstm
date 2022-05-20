package config

import (
	"os"
	"reflect"

	"github.com/spf13/viper"
)

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	// Use reflect to extend environment variables for each field
	pathStruct := reflect.ValueOf(&conf.Path).Elem()
	for i := 0; i < pathStruct.NumField(); i++ {
		oldVal := pathStruct.Field(i).String()
		pathStruct.Field(i).SetString(os.ExpandEnv(oldVal))
	}

	return conf
}
