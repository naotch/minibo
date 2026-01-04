package config

import (
	"os"
)

var Config Values

type Values struct {
	Env    string
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
}

func init() {
	Config = Values{
		Env:    getEnv("APP_ENV", "dev"),
		DBHost: getEnv("DB_HOST", "0.0.0.0"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBName: getEnv("DB_NAME", "postgres"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASS", "postgres"),
	}
}

func getEnv(key, def string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return def
	}
	return val
}
