package config

import (
	"path/filepath"
	"runtime"

	"github.com/josephpballantyne/go-project-template/internal/app"
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
	const op = "config.InitViper"
	viper.SetConfigName("config")
	viper.AddConfigPath(Root)
	err := viper.ReadInConfig()
	if err != nil {
		return Constants{}, &app.Error{Op: op, Err: err}
	}
	viper.SetDefault("PORT", "3000")

	var constants Constants
	err = viper.Unmarshal(&constants)
	if err != nil {
		return Constants{}, &app.Error{Op: op, Err: err}
	}
	return constants, nil
}
