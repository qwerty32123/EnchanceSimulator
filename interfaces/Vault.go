package interfaces

type VaultService interface {
	GetSecret(path string, key string) (string, error)
}
