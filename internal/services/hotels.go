package services

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
	"githumb/go-related/nuitteassignment/internal/models"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	statusUrl = "status"
	hotelUrl  = "hotels"
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

func (s *Service) CheckHotelRate(checkinTime, checkoutTime *time.Time, currency, guestNationality string, hotelIds []uint, occupancies []models.Occupancies) (bool, error) {
	fmt.Println("test")
	return true, nil
}

func NewService(apiKey, apiSecret, baseUrl string) (*Service, error) {
	return &Service{apiKey, apiSecret, baseUrl}, nil
}

func (s *Service) prepareHttpRequest(targetPath string) (*http.Request, error) {
	// calculate the signature
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	signatureString := s.ApiKey + s.ApiSecret + timestamp
	hasher := sha256.New()
	hasher.Write([]byte(signatureString))
	signatureEncoded := hex.EncodeToString(hasher.Sum(nil))

	// prepare the request with appropriate headers
	url := fmt.Sprintf("%s/%s", s.BaseURL, targetPath)
	req, err := http.NewRequest(http.MethodGet, url, nil)
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

func (s *Service) CheckStatus() (bool, error) {

	// prepare the request
	req, err := s.prepareHttpRequest(statusUrl)
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