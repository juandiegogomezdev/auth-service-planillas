package config

type ConfigVar struct {
	JWTSecret string
}

func LoadConfig() ConfigVar {
	return ConfigVar{
		JWTSecret: GetEnv("JWT_SECRET", "supersecretkey"),
	}
}

var Config = LoadConfig()
