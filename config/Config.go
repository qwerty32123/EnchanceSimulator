package config

import (
	"encoding/json"
	"os"
	"sync"
)

type APIConfigurator interface {
	GetAPIURL() string
	GetRegions() map[string]string
	GetEndpoints() map[string]string
	GetLanguages() []string
	GetRabbitMQURL() string
	GetUserAgent() string
	GetContentType() string
}

// JSONConfig represents the JSON configuration structure.
type JSONConfig struct {
	API         string            `json:"API"`
	Regions     map[string]string `json:"REGION"`
	Languages   []string          `json:"LANGUAGE"`
	Endpoints   map[string]string `json:"ENDPOINTS"`
	RabbitMQURL string            `json:"RabbitMQURL"`
	VaultAPIURL string            `json:"VaultAPIURL"`
	VaultToken  string            `json:"VaultToken"`
	UserAgent   string            `json:"UserAgent"`
	ContentType string            `json:"ContentType"`
	InputQueu   string            `json:"InputQueu"`
	OutputQueu  string            `json:"InputQueu"`
}

var appConfig *JSONConfig
var configOnce sync.Once

// LoadConfigFromJSONFile loads the JSON configuration from a file.
func LoadConfigFromJSONFile(filePath string) error {
	var err error
	configOnce.Do(func() {
		data, readFileErr := os.ReadFile(filePath)
		if readFileErr != nil {
			err = readFileErr
			return
		}

		err = json.Unmarshal(data, &appConfig)
	})

	return err
}

// GetConfig returns the loaded JSON configuration.
func GetConfig() *JSONConfig {
	return appConfig
}

func (jc *JSONConfig) GetAPIURL() string {
	return jc.API
}
func (jc *JSONConfig) GetVaultApiUrl() string {
	return jc.VaultAPIURL
}
func (jc *JSONConfig) GetVaultToken() string {
	return jc.VaultToken
}

func (jc *JSONConfig) GetRegions() map[string]string {
	return jc.Regions
}

func (jc *JSONConfig) GetEndpoints() map[string]string {
	return jc.Endpoints
}

func (jc *JSONConfig) GetLanguages() []string {
	return jc.Languages
}

func (jc *JSONConfig) GetRabbitMQURL() string {
	return jc.RabbitMQURL
}

func (jc *JSONConfig) GetUserAgent() string {
	return jc.UserAgent
}

func (jc *JSONConfig) GetContentType() string {
	return jc.ContentType
}
