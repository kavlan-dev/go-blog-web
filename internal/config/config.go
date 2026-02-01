package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Env   string // Окружение может быть local, dev, prod
	Admin struct {
		Username string
		Password string
		Email    string
	}
	Server struct {
		Host string
		Port uint
	}
	CORSAllowedOrigin []string
}

func InitConfig() (*Config, error) {
	var config Config
	config.Env = envOrDefault("ENV", "prod")
	config.Admin.Username = envOrDefault("ADMIN_USERNAME", "admin")
	config.Admin.Password = envOrDefault("ADMIN_PASSWORD", "admin")
	config.Admin.Email = envOrDefault("ADMIN_EMAIL", "admin@example.com")
	config.Server.Host = envOrDefault("SERVER_HOST", "localhost")
	config.CORSAllowedOrigin = strings.Split(envOrDefault("CORS_ALLOWED_ORIGIN", "*"), ",")

	port, err := strconv.Atoi(envOrDefault("SERVER_PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("не верное значение порта: %s", err.Error())
	}
	config.Server.Port = uint(port)

	if config.Env != "local" && config.Env != "dev" && config.Env != "prod" {
		return nil, fmt.Errorf("не верное значение окружения: %s", config.Env)
	}

	return &config, nil
}

// Возвращает адрес на котором запускается сервер
func (c *Config) ServerAddress() string {
	port := strconv.Itoa(int(c.Server.Port))
	return c.Server.Host + ":" + port
}

func (c *Config) Cors() string {
	return strings.Join(c.CORSAllowedOrigin, ", ")
}

func envOrDefault(varName string, defaultValue string) string {
	value := os.Getenv(varName)
	if value == "" {
		value = defaultValue
	}

	return value
}
