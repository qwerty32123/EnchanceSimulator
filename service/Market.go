package services

import (
	"EnchanceSimulator/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Market defines the interface that your service should implement
type Market interface {
	MakeHTTPPostRequest() ([]byte, error)
}
type LocalMarketService struct {
	models.MarketServiceImpl
}

// NewMarketService is the constructor function for MarketServiceImpl
func NewMarketService(config models.MarketConfig) *models.MarketServiceImpl {
	return &models.MarketServiceImpl{
		Config: config,
	}
}
func (s LocalMarketService) MakeHTTPPostRequest() ([]byte, error) {
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	return responseBytes, nil
}
