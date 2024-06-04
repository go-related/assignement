package configurations

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AssignmentConfigurations struct {
	BaseUrl   string
	Port      string
	ApiSecret string
	ApiKey    string
	LogLevel  string
}

func NewAssignmentConfigurations() (*AssignmentConfigurations, error) {
	v := viper.New()
	v.SetConfigName("nuitee")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file because we
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.Is(err, &configFileNotFoundError) {
			logrus.WithError(err).Warning("error loading config file")
		}
	}

	var config AssignmentConfigurations
	err := v.UnmarshalExact(&config)
	return &config, err
}
