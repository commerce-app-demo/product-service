// internal/config/config.go
package config

import (
	"os"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type ServerConfig struct {
	Port string
	Host string
}

func Load() *Config {
	return &Config{
		Database: DatabaseConfig{
			Driver:   "mysql",
			Host:     getEnvOrDefault("DB_HOST", "localhost"),
			Port:     getEnvOrDefault("DB_PORT", "3306"),
			User:     getEnvOrDefault("DB_USER", "product_user"),
			Password: getEnvOrDefault("DB_PASSWORD", "product_pass"),
			DBName:   getEnvOrDefault("DB_NAME", "product_db"),
		},
		Server: ServerConfig{
			Port: getEnvOrDefault("SERVER_PORT", "50051"),
			Host: getEnvOrDefault("SERVER_HOST", "localhost"),
		},
	}
}

func LoadDBConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Driver:   "mysql",
		Host:     getEnvOrDefault("DB_HOST", "localhost"),
		Port:     getEnvOrDefault("DB_PORT", "3306"),
		User:     getEnvOrDefault("DB_USER", "product_user"),
		Password: getEnvOrDefault("DB_PASSWORD", "product_pass"),
		DBName:   getEnvOrDefault("DB_NAME", "product_db"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
