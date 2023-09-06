package services

import (
	"EnchanceSimulator/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type FanceApiResponse struct {
	Data []models.FancyResponseApi `json:"data"`
}

func FetchItemsAsync(filter string, ch chan []models.FancyResponseApi) {
	url := fmt.Sprintf(
		"https://secret/en/EU/db/items?page=1&filter=%s",
		filter,
	)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		ch <- nil
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		ch <- nil
		return
	}

	var FancyResp FanceApiResponse
	err = json.Unmarshal(body, &FancyResp)
	if err != nil {
		fmt.Println("Error:", err)
		ch <- nil
		return
	}

	ch <- FancyResp.Data
}
