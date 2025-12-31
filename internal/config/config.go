package config

import (
	"fmt"
	"os"
)

// Config holds all application configuration
type Config struct {
	ServiceName string
	DBDriver    string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	FrontendURL string 

}

// LoadConfig reads environment variables
func LoadConfig() (*Config, error) {
	cfg := &Config{
		ServiceName: getEnv("SERVICE_NAME", "user.service"),
		DBDriver:    getEnv("DB_DRIVER", "pgx"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "user_service"),
		DBPassword:  getEnv("DB_PASSWORD", "password"),
		DBName:      getEnv("DB_NAME", "userdb"),
		FrontendURL: getEnv("FRONTEND_URL", "http://localhost:5173"),
	}

	return cfg, nil
}

// Helper: get env with default value
func getEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

// Build PostgreSQL DSN
func (c *Config) DatabaseDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}
