package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() Config {
	// Load .env file if it exists (ignore errors - file might not exist)
	_ = godotenv.Load()

	return Config{
		Environment: EnvironmentConfig{
			Name: getEnvWithDefault("ENVIRONMENT_NAME", "development"),
		},
		Server: ServerConfig{
			Port:     getEnvWithDefault("SERVER_PORT", "8080"),
			LogLevel: getEnvWithDefault("SERVER_LOG_LEVEL", "info"),
		},
		Database: DatabaseConfig{
			Host:     getEnvWithDefault("DATABASE_HOST", "localhost"),
			Port:     getEnvWithDefault("DATABASE_PORT", "5432"),
			Name:     getEnvWithDefault("DATABASE_NAME", "template"),
			User:     getEnvWithDefault("DATABASE_USER", "template"),
			Password: getEnvWithDefault("DATABASE_PASSWORD", "template"),
			SSLMode:  getEnvWithDefault("DATABASE_SSL_MODE", "disable"),
		},
		Session: SessionConfig{
			Secret: getEnvWithDefault("SESSION_SECRET", "cookie_secret_key"),
			Domain: getEnvWithDefault("SESSION_DOMAIN", "localhost"),
		},
		Frontend: FrontendConfig{
			URL: getEnvWithDefault("FRONTEND_URL", "http://localhost:5173"),
		},
	}
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
