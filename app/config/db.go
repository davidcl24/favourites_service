// Configuration module that lets you set up the database connection
package config

import "os"

// The struct that contains all data for the database connection
type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

// Creates a new DBConfig struct. It gets the data from environment variables.
// If they happened to not exist, it will take a fallback value.
func NewEnvDBConfig() *DBConfig {
	return &DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		Username: getEnv("DB_USERNAME", "default"),
		Password: getEnv("DB_PASSWORD", "example"),
		Database: getEnv("DB_DATABASE", "streamingdb"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
