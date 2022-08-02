package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

const (
	redis_database_url = "REDIS_URL"
	redis_host         = "REDIS_HOST"
	redis_port         = "REDIS_PORT"
)

func init() {
	if err := godotenv.Load("./config/env/.env"); err != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", err))
	}
}

type RedisConfig struct {
	RedisUrl string
	Host     string
	Port     string
}

func GetRedisConfig() *RedisConfig {
	return &RedisConfig{
		os.Getenv(redis_database_url),
		os.Getenv(redis_host),
		os.Getenv(redis_port),
	}
}
