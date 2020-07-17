package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const DefaultEnv = "local"

func LoadConfig() {
	_, p, _, _ := runtime.Caller(0)
	base := filepath.Dir(p)

	env := strings.ToLower(os.Getenv("ENV"))
	if len(env) == 0 {
		env = DefaultEnv
	}
	viper.Set("env", env)

	configFileName := fmt.Sprintf("config-%v", env)

	logrus.Infof("Load config file: %v", configFileName)
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(base)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("config file is not found: %v", err))
		} else {
			panic(fmt.Errorf("config file is found but got error: %v", err))
		}
	}
}
