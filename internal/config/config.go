package config

import "github.com/spf13/viper"

type Constants struct {
	PORT  string
	Mongo struct {
		URL    string
		DBName string
	}
}

func InitViper() (Constants, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return Constants{}, err
	}
	viper.SetDefault("PORT", "3000")

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}
