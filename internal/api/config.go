package api

import (
	"fmt"
	"github.com/fumiyauehara/go-api/internal/api/model"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func InitConfig(envPath string) model.Config {
	err := godotenv.Load(envPath)
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %s", err))
	}

	return model.Config{
		Postgres: model.DBConfig{
			Type:   "postgres",
			DBHost: os.Getenv("POSTGRES_HOST"),
			DBPort: convertStringToInt(os.Getenv("POSTGRES_PORT")),
			DBName: os.Getenv("POSTGRES_DB"),
			DBUser: os.Getenv("POSTGRES_USER"),
			DBPw:   os.Getenv("POSTGRES_PW"),
		},
		MysqlWriter: model.DBConfig{
			Type:   "mysql",
			DBHost: os.Getenv("MYSQL_WRITE_HOST"),
			DBPort: convertStringToInt(os.Getenv("MYSQL_WRITE_PORT")),
			DBName: os.Getenv("MYSQL_WRITE_DB"),
			DBUser: os.Getenv("MYSQL_WRITE_USER"),
			DBPw:   os.Getenv("MYSQL_WRITE_PW"),
		},
		MysqlReader: model.DBConfig{
			Type:   "mysql",
			DBHost: os.Getenv("MYSQL_READ_HOST"),
			DBPort: convertStringToInt(os.Getenv("MYSQL_READ_PORT")),
			DBName: os.Getenv("MYSQL_READ_DB"),
			DBUser: os.Getenv("MYSQL_READ_USER"),
			DBPw:   os.Getenv("MYSQL_READ_PW"),
		},
		RedisConfig: model.RedisConfig{
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
