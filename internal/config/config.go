package config

import (
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../..")
)

type Constants struct {
	PORT  string
	Mongo struct {
		URL    string
		DBName string
	}
}

func InitViper() (Constants, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(Root)
	err := viper.ReadInConfig()
	if err != nil {
		return Constants{}, err
	}
	viper.SetDefault("PORT", "3000")

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}
