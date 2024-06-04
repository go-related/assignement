package services

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"githumb/go-related/nuitteassignment/internal/models"
	"githumb/go-related/nuitteassignment/internal/models/response"
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
	CheckStatus() (bool, []byte, []byte, error)
	CheckHotelRate(checkinTime, checkoutTime *time.Time, currency, guestNationality string, hotelIds []uint, occupancies []models.Occupancies) ([]models.HotelResponse, []byte, []byte, error)
}

type Service struct {
	ApiKey    string
	ApiSecret string
	BaseURL   string
}

func NewService(apiKey, apiSecret, baseUrl string) (*Service, error) {
	return &Service{apiKey, apiSecret, baseUrl}, nil
}

func (s *Service) CheckHotelRate(checkinTime, checkoutTime *time.Time, currency, guestNationality string, hotelIds []uint, occupancies []models.Occupancies) ([]models.HotelResponse, []byte, []byte, error) {
	var output []models.HotelResponse
	// check the validation for simplicity we will make everything required
	if utils.IsDateEmpty(checkinTime) || utils.IsDateEmpty(checkoutTime) {
		return output, nil, nil, NewServiceError("invalid data for checkin-time or checkout-time")
	}
	// currency and guestNationality is not required on the

	if len(hotelIds) == 0 {
		return output, nil, nil, NewServiceError("invalid hotel-ids")
	}
	if len(occupancies) == 0 {
		return output, nil, nil, NewServiceError("invalid occupancies")
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
	req, requestBody, err := s.prepareHttpRequest(http.MethodPost, hotelsUrl, requestData)
	if err != nil {
		return output, requestBody, nil, err
	}
	_, responseBody, err := s.doHttpCall(req)
	if err != nil {
		logrus.WithError(err).Error("failed to check hotel rates")
		return output, requestBody, responseBody, err
	}
	var response response.HotelResponseDTO
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		logrus.WithError(err).Error("failed to read hotel response")
		return output, requestBody, responseBody, err
	}
	if response.Hotels != nil && response.Hotels.Hotels != nil {
		for _, hotel := range *response.Hotels.Hotels {
			output = append(output, models.HotelResponse{
				HotelId:  hotel.Code,
				Currency: hotel.Currency,
				Price:    hotel.MinRate,
			})
		}
	}

	return output, requestBody, responseBody, nil
}

func (s *Service) CheckStatus() (bool, []byte, []byte, error) {

	// prepare the request
	req, requestBody, err := s.prepareHttpRequest(http.MethodGet, statusUrl, nil)
	if err != nil {
		return false, requestBody, nil, err
	}

	status, body, err := s.doHttpCall(req)
	if err != nil {
		logrus.WithError(err).Error("failed to check hotel status")
		return false, requestBody, body, err
	}

	// this is a simplified version of checking this
	if status == 200 && strings.Contains(string(body), `"status":"OK"`) {
		return true, requestBody, body, nil
	}

	return false, requestBody, body, nil
}

func (s *Service) prepareHttpRequest(method, targetUrl string, bodyData interface{}) (*http.Request, []byte, error) {
	var requestBody []byte
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
			return nil, requestBody, err
		}
		body = bytes.NewBuffer(jsonBody)
		requestBody = jsonBody
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, requestBody, err
	}
	req.Header.Set("Accept", "application/json") // all these need to constants
	req.Header.Set("Api-key", s.ApiKey)
	req.Header.Set("X-Signature", signatureEncoded)
	if bodyData != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, requestBody, nil
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
