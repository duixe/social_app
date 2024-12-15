package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
	CurrentEnv             string
	Version                string
}

// create a singleton to be used in the app
var Envs = initConfig()

func initConfig() AppConfig {
	godotenv.Load()

	return AppConfig{
		PublicHost:             GetString("PUBLIC_HOST", "http://localhost"),
		Port:                   GetString("PORT", ":8080"),
		DBUser:                 GetString("DB_USER", "root"),
		DBPassword:             GetString("DB_PASSWORD", ""),
		DBAddress:              GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/gosocial?sslmode=disable"),
		DBName:                 GetString("DB_NAME", "go_rest_first"),
		JWTExpirationInSeconds: GetInt("JWT_EXP", 3600*24*7),
		JWTSecret:              GetString("JWT_SECRET", "it-just-a-secret-duh:)"),
		CurrentEnv:             GetString("ENV", "development"),
		Version:                GetString("APP_VERSION", "1.0.0"),
	}
}

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
