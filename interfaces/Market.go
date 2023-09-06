package interfaces

type MarketService interface {
	MakeHTTPPostRequest() ([]byte, error)
}
