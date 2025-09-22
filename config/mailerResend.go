package config

import "os"

type resendConfig struct {
	APIKey string
}

var Resend = resendConfig{
	APIKey: os.Getenv("RESEND_API_KEY"),
}
