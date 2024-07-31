package config

import (
	"log"
	"os"
	"strconv"
)

// Config struct now directly reads from environment variables
type Config struct {
	DB            DBConfig
	UserService   ServiceConfig `env:"USER_SERVICE_"`
	AuthService   ServiceConfig `env:"AUTH_SERVICE_"`
	JWT           JWTConfig
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type ServiceConfig struct {
	Host string
	Port int
}

type JWTConfig struct {
	Secret string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	cfg := &Config{
		DB: DBConfig{
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", ""),
		},
		UserService: ServiceConfig{
			Host: getEnv("USER_SERVICE_HOST", ""),
			Port: getEnvAsInt("USER_SERVICE_PORT", 0),
		},
		AuthService: ServiceConfig{
			Host: getEnv("AUTH_SERVICE_HOST", ""),
			Port: getEnvAsInt("AUTH_SERVICE_PORT", 8080),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", ""),
		},
	}

	log.Println("Config loaded successfully")
	return cfg, nil
}

// getEnv gets an environment variable or returns a default value if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt gets an environment variable as an integer or returns a default value if not set
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
