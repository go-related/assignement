package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
	"githumb/go-related/nuitteassignment/internal"
	"githumb/go-related/nuitteassignment/internal/configurations"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// API KEY:    db11033c50b5ed53ab7b815cb1b2eaee
// API SECRET:  704773c0e3
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

func testApi() {
	apiKey := "db11033c50b5ed53ab7b815cb1b2eaee"
	secret := "704773c0e3"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// Concatenate apiKey, secret, and timestamp
	signatureString := apiKey + secret + timestamp

	// Generate SHA-256 hash
	hasher := sha256.New()
	hasher.Write([]byte(signatureString))
	xSignature := hex.EncodeToString(hasher.Sum(nil))

	// Create a new request
	url := "https://api.test.hotelbeds.com/hotel-api/1.0/status"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	// Set the headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Api-key", apiKey)
	req.Header.Set("X-Signature", xSignature)

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Println("Response Status:", resp.Status)

	// Print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}
	fmt.Println("Response Body:", string(body))
}
