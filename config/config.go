package config

import (
	"path/filepath"
	"strings"

	"github.com/PlanckProject/go-commons/logger"
	"github.com/spf13/viper"
)

func Parse(config interface{}, path string) {
	v := viper.New()

	logger.Infof("Config path: %s", path)
	configDir := filepath.Dir(path)
	filename := filepath.Base(path)

	filenameTokens := strings.Split(filename, ".")

	v.AddConfigPath(configDir)
	v.SetConfigName(filenameTokens[0])
	v.SetConfigType(filenameTokens[1])

	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Fatalln("Config not found in directory")
		} else {
			logger.Fatalln(err)
		}
	}
	err = v.Unmarshal(config)
	if err != nil {
		logger.Fatalln(err)
	}
}
