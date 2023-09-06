package services

import (
	"EnchanceSimulator/config"
	"fmt"
	vault "github.com/hashicorp/vault/api"
)

// VaultService interface for interacting with Vault
type VaultService interface {
	GetSecret(path string, key string) (string, error)
}

// Secret struct to hold key-value pair
type Secret struct {
	Key   string
	Value string
}

// VaultServiceImpl implementation of VaultService
type VaultServiceImpl struct {
	client *vault.Client
}

// NewVaultService function to initialize Vault service
func NewVaultService(config *config.JSONConfig) (VaultService, error) {
	clientConfig := &vault.Config{
		Address: config.GetAPIURL(),
	}

	client, err := vault.NewClient(clientConfig)
	if err != nil {
		return nil, err
	}

	// Assumes Vault token is stored in the "UserAgent" field of JSONConfig. Please secure this properly in production.
	client.SetToken(config.GetVaultToken())

	return &VaultServiceImpl{
		client: client,
	}, nil
}

// GetSecret method to fetch a secret from Vault
func (s *VaultServiceImpl) GetSecret(path string, key string) (string, error) {
	secret, err := s.client.Logical().Read(path)
	if err != nil {
		return "", err
	}

	if secret == nil || secret.Data == nil {
		return "", fmt.Errorf("secret not found")
	}

	data := secret.Data["data"].(map[string]interface{})
	value, ok := data[key].(string)
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	return value, nil
}
