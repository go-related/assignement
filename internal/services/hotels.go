package services

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"githumb/go-related/nuitteassignment/internal/models"
	"githumb/go-related/nuitteassignment/internal/utils"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	statusUrl = "status"
	hotelsUrl = "hotels"
)

type Hotel interface {
	CheckStatus() (bool, error)
	CheckHotelRate(checkinTime, checkoutTime *time.Time, currency, guestNationality string, hotelIds []uint, occupancies []models.Occupancies) (bool, error)
}

type Service struct {
	ApiKey    string
	ApiSecret string
	BaseURL   string
}

func NewService(apiKey, apiSecret, baseUrl string) (*Service, error) {
	return &Service{apiKey, apiSecret, baseUrl}, nil
}

func (s *Service) CheckHotelRate(checkinTime, checkoutTime *time.Time, currency, guestNationality string, hotelIds []uint, occupancies []models.Occupancies) (bool, error) {
	// check the validation for simplicity we will make everything required
	if utils.IsDateEmpty(checkinTime) || utils.IsDateEmpty(checkoutTime) {
		return false, NewServiceError("invalid data for checkin-time or checkout-time")
	}
	// currency and guestNationality is not required on the

	if len(hotelIds) == 0 {
		return false, NewServiceError("invalid hotel-ids")
	}
	if len(occupancies) == 0 {
		return false, NewServiceError("invalid occupancies")
	}

	requestData := models.HotelRequestDTO{
		Stay: models.StayDTO{
			CheckIn:  utils.ConvertDateToString(*checkinTime),
			CheckOut: utils.ConvertDateToString(*checkoutTime),
		},
		Hotels: models.HotelDTOs{
			Hotel: hotelIds,
		},
		Occupancies: convertOccupancies(occupancies),
	}
	req, err := s.prepareHttpRequest(http.MethodPost, hotelsUrl, requestData)
	if err != nil {
		return false, err
	}
	status, responsebody, err := s.doHttpCall(req)
	if err != nil {
		logrus.WithError(err).Error("failed to check hotel rates")
		return false, err
	}
	println(status)
	println(string(responsebody))
	return true, nil
}

func (s *Service) CheckStatus() (bool, error) {

	// prepare the request
	req, err := s.prepareHttpRequest(http.MethodGet, statusUrl, nil)
	if err != nil {
		return false, err
	}

	status, body, err := s.doHttpCall(req)
	if err != nil {
		logrus.WithError(err).Error("failed to check hotel status")
		return false, err
	}

	// this is a simplified version of checking this
	if status == 200 && strings.Contains(string(body), `"status":"OK"`) {
		return true, nil
	}

	return false, nil
}

func (s *Service) prepareHttpRequest(method, targetUrl string, bodyData interface{}) (*http.Request, error) {
	// calculate the signature
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	signatureString := s.ApiKey + s.ApiSecret + timestamp
	hasher := sha256.New()
	hasher.Write([]byte(signatureString))
	signatureEncoded := hex.EncodeToString(hasher.Sum(nil))

	// prepare the request with appropriate headers
	url := fmt.Sprintf("%s/%s", s.BaseURL, targetUrl)
	var body io.Reader
	if bodyData != nil {
		jsonBody, err := json.Marshal(bodyData)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonBody)
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json") // all these need to constants
	req.Header.Set("Api-key", s.ApiKey)
	req.Header.Set("X-Signature", signatureEncoded)
	return req, nil
}

func (s *Service) doHttpCall(req *http.Request) (int, []byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.WithError(err).Error("failed to check hotel status")
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err).Error("failed to read response body")
	}
	return resp.StatusCode, body, nil
}
