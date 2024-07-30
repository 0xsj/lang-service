package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	DB            DBConfig
	UserService   ServiceConfig `mapstructure:"user-service"`
	AuthService   ServiceConfig `mapstructure:"auth-service"`
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
	Port string
}

type JWTConfig struct {
	Secret string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	overrideConfigWithEnv(&cfg)

	log.Println("Config loaded successfully")
	return &cfg, nil
}

func overrideConfigWithEnv(cfg *Config) {
	if portStr := os.Getenv("PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			cfg.DB.Port = port
		}
	}

	if host := os.Getenv("HOST"); host != "" {
		cfg.DB.Host = host
	}
}
