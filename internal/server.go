package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"githumb/go-related/nuitteassignment/internal/configurations"
	"githumb/go-related/nuitteassignment/internal/handlers"
	"githumb/go-related/nuitteassignment/internal/services"
)

type Server struct {
	handler *handlers.Handler
}

func NewServer(config *configurations.AssignmentConfigurations) (*Server, error) {

	//setup dependencies
	router := gin.Default()

	hotelService, err := services.NewHotelService(config.ApiKey, config.ApiSecret, config.BaseUrl)
	if err != nil {
		return nil, err
	}
	// setup routes
	handler := handlers.NewHandler(hotelService, router)

	err = router.Run(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		logrus.WithError(err).Errorf("Setting up service failed.")
		return nil, err
	}
	logrus.Infof("Application is running on port:%s", config.Port)
	return &Server{handler: handler}, nil
}
