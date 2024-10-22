package config

type Config struct {
	DatabaseURL string
}

var config *Config

func New() *Config {
	if config != nil {
		return config
	}

	config = &Config{
		DatabaseURL: databaseUrl(),
	}

	return config
}
