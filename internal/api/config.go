package api

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type DBConfig struct {
	DBHost string
	DBPort int
	DBName string
	DBUser string
	DBPass string
}
type RedisConfig struct {
	RedisHost string
	RedisPort int
}

type Config struct {
	DBConfig
	RedisConfig
	Port int
}

func InitConfig(envPath string) Config {
	err := godotenv.Load(envPath)
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %s", err))
	}

	return Config{
		DBConfig: DBConfig{
			DBHost: os.Getenv("POSTGRES_HOST"),
			DBPort: convertStringToInt(os.Getenv("POSTGRES_PORT")),
			DBName: os.Getenv("POSTGRES_DB"),
			DBUser: os.Getenv("POSTGRES_USER"),
			DBPass: os.Getenv("POSTGRES_PASSWORD"),
		},
		RedisConfig: RedisConfig{
			RedisHost: os.Getenv("REDIS_HOST"),
			RedisPort: convertStringToInt(os.Getenv("REDIS_PORT")),
		},
		Port: convertStringToInt(os.Getenv("APP_PORT")),
	}
}

func convertStringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Error converting %s to int from setting env: %s", s, err))
	}
	return num
}
