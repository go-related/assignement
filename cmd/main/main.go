package main

import (
	"github.com/sirupsen/logrus"
	"githumb/go-related/nuitteassignment/internal"
	"githumb/go-related/nuitteassignment/internal/configurations"
)

func main() {
	configs, err := configurations.NewAssignmentConfigurations()
	if err != nil {
		logrus.WithError(err).Error("failed to load configurations.")
	}
	_, err = internal.NewServer(configs)
	if err != nil {
		logrus.WithError(err).Error("failed to setup server.")
	}
}
