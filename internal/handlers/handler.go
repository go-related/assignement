package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"githumb/go-related/nuitteassignment/internal/services"
	"githumb/go-related/nuitteassignment/internal/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Handler struct {
	Service services.Hotel
	Engine  *gin.Engine
}

type SupplierData struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}
type Response struct {
	Data     interface{}  `json:"data"`
	Err      string       `json:"error_message"`
	Supplier SupplierData `json:"supplier"`
}

func NewHandler(hotelService services.Hotel, router *gin.Engine) *Handler {
	handler := &Handler{
		Service: hotelService,
		Engine:  router,
	}
	router.GET("/status", handler.GetHealth)
	v1 := router.Group("/api/v1")

	// register country
	v1.GET("/hotels", handler.GetHotels)

	return handler
}

// helpers
func getQueryParamDate(c *gin.Context, paramName string) (time.Time, error) {
	id := c.Query(paramName)
	return utils.ConvertStringToDate(id)
}

func getQueryParamIdList(c *gin.Context, paramName string) ([]uint, error) {
	id := c.Query(paramName)
	if id != "" {
		values := strings.Split(id, ",")
		if len(values) > 0 {
			output := make([]uint, len(values))
			for counter, v := range values {
				idValue, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					return nil, err
				}
				output[counter] = uint(idValue)
			}
			return output, nil
		}
	}
	return []uint{}, nil
}

func getQueryParamDataList[T any](c *gin.Context, paramName string) ([]T, error) {
	value := c.Query(paramName)
	if value != "" {
		var result []T
		err := json.Unmarshal([]byte(value), &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, fmt.Errorf("qyery parameter not found")
}

func returnOk(c *gin.Context, status int, data interface{}, request, response []byte) {
	c.IndentedJSON(status, Response{
		Data: data,
		Supplier: SupplierData{
			Request:  string(request),
			Response: string(response),
		},
	})
}

func AbortWithMessage(c *gin.Context, status int, err error, message string, request, response []byte) {
	logrus.WithError(err).Error(message)

	// custom validation error update status and message
	var badRequest *services.ServiceError
	if errors.As(err, &badRequest) {
		status = http.StatusBadRequest
		message = err.Error()
	}
	c.AbortWithStatusJSON(status, Response{
		Err: errors.New(message).Error(), // so we don't send stack trace to the clients
		Supplier: SupplierData{
			Request:  string(request),
			Response: string(response),
		},
	})
}
