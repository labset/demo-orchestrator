package config

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

type Config struct {
	Environment EnvironmentConfig
	Server      ServerConfig
	Database    DatabaseConfig
	Session     SessionConfig
	Frontend    FrontendConfig
}

type EnvironmentConfig struct {
	Name string
}

func (e *EnvironmentConfig) IsDevelopment() bool {
	env := strings.ToLower(e.Name)

	return env == "development" || env == "dev"
}

func (e *EnvironmentConfig) IsProduction() bool {
	env := strings.ToLower(e.Name)

	return env == "production" || env == "prod"
}

func (e *EnvironmentConfig) IsStaging() bool {
	env := strings.ToLower(e.Name)

	return env == "staging" || env == "stage"
}

type ServerConfig struct {
	LogLevel string
	Port     string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}

func (d *DatabaseConfig) ConnectionString() string {
	target := net.JoinHostPort(d.Host, d.Port)
	credentials := url.UserPassword(d.User, d.Password).String()

	return fmt.Sprintf(
		"postgres://%s@%s/%s?sslmode=%s",
		credentials, target,
		d.Name,
		d.SSLMode,
	)
}

type SessionConfig struct {
	Secret string
	Domain string
}

type FrontendConfig struct {
	URL string
}
