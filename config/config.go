package config

import "github.com/339-Labs/v3-bitget-api-sdk-go/constants"

type BitgetConfig struct {
	BaseUrl       string
	WsUrl         string
	ApiKey        string
	SecretKey     string
	PASSPHRASE    string
	TimeoutSecond int
	SignType      string // 可选: "HMAC_SHA256" or "RSA"
}

func NewBitgetConfig(ApiKey string, SecretKey string, PASSPHRASE string, TimeoutSecond int, SignType string) *BitgetConfig {
	if SignType == "" {
		SignType = constants.SHA256
	}
	return &BitgetConfig{
		BaseUrl:       "https://api.bitget.com",
		WsUrl:         "wss://ws.bitget.com/mix/v1/stream",
		ApiKey:        ApiKey,
		SecretKey:     SecretKey,
		PASSPHRASE:    PASSPHRASE,
		TimeoutSecond: TimeoutSecond,
		SignType:      SignType,
	}
}
