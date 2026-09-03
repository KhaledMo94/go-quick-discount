package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func getEnvOrDefault(key string , defaultValue string) string{
	env := os.Getenv(key)
	if env == "" {
		return  defaultValue
	}
	return env
}

type Config struct {
	DBUser     	string
	DBPassword 	string
	DBHost     	string
	DBPort     	string
	DBName     	string
	Driver		string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	MEILISEARCH_HOST string
	MEILISEARCH_KEY string
	SCOUT_PREFIX string
}

func Load() (*Config , error){

	if err := godotenv.Load(); err != nil{
		slog.Debug(".env file not found, relying on system env vars( if applicable)")
	}else{
		slog.Debug("loaded configuration from .env file")
	}

	redisDB, err := strconv.Atoi(getEnvOrDefault("REDIS_DB", "0"))
	if err != nil {
		return nil, fmt.Errorf("load config: invalid REDIS_DB: %w", err)
	}

	cfg := &Config{
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_CONNECTION"),
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_PASSWORD"),
		redisDB,
		os.Getenv("MEILISEARCH_HOST"),
		os.Getenv("MEILISEARCH_KEY"),
		os.Getenv("SCOUT_PREFIX"),
	}

	if cfg.DBName == "" || cfg.DBUser == "" || cfg.DBHost == "" || cfg.DBPort == "" {
		return nil , fmt.Errorf("load config: missing required DB_USER or DB_NAME")
	}

	return  cfg , nil
}

func (cfg *Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
}

func (cfg *Config) RedisAddr() string {
	return fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
}

