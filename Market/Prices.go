package Market

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// MakeHTTPPostRequest sends an HTTP POST request and returns the response bytes.
func MakeHTTPPostRequest() ([]byte, error) {
	// Define the request payload for the POST request
	payload := map[string]interface{}{
		"keytype":      0,
		"maincategory": 10,
		"subcategory":  3,
	}

	// Convert the payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("Error encoding JSON: %v", err)
	}

	// Define the URL for the POST request to XXX servers
	url := "secret/Trademarket/GetWorldMarketHotList"

	// Create an HTTP client
	client := &http.Client{}

	// Create the POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %v", err)
	}

	// Set headers for the request
	req.Header.Set("User-Agent", "XXX")
	req.Header.Set("Content-Type", "application/json")

	// Make the POST
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error making request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Read and return the response bytes
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %v", err)
	}

	return responseBytes, nil
}

func main() {
	// Make the HTTP POST request and get the response bytes
	responseBytes, err := MakeHTTPPostRequest()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response from  servers:")
	fmt.Println(string(responseBytes))

	// Continue with the rest of your code...
}
