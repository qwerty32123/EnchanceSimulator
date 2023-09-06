package services

import (
	"BdoEnchanceApi/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// BdoService defines the interface that your service should implement
type BdoService interface {
	MakeHTTPPostRequest() ([]byte, error)
}
type LocalBdoMarketService struct {
	models.BdoMarketServiceImpl
}

// NewBdoMarketService is the constructor function for BdoMarketServiceImpl
func NewBdoMarketService(config models.BdoMarketConfig) *models.BdoMarketServiceImpl {
	return &models.BdoMarketServiceImpl{
		Config: config,
	}
}
func (s LocalBdoMarketService) MakeHTTPPostRequest() ([]byte, error) {
	payload := map[string]interface{}{
		"keytype":      0,
		"maincategory": 10,
		"subcategory":  3,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error encoding JSON: %v", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", s.Config.URL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("User-Agent", s.Config.UserAgent)
	req.Header.Set("Content-Type", s.Config.ContentType)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	return responseBytes, nil
}
