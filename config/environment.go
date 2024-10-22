package config

import "os"

func databaseUrl() string {
	return os.Getenv("DB_URL")
}
