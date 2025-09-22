package config

import "os"

type jwtConfig struct {
	JWTSecret string
}

var JWT = jwtConfig{
	JWTSecret: os.Getenv("JWT_SECRET"),
}
