package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Use LoadConfig to create one.
type Config struct{}

// LoadConfig loads the given .env file (or the default .env when path == "")
// and returns a Config instance.
func LoadConfig(path string) *Config {
	if path == "" {
		_ = godotenv.Load()
	} else {
		_ = godotenv.Load(path)
	}
	return &Config{}
}

// Get returns the environment variable value for key, or the provided default.
func (c *Config) Get(key string, defaultValue ...string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

// GetInt returns the environment variable parsed as int, or the provided default.
func (c *Config) GetInt(key string, defaultValue ...int) int {
	s := c.Get(key)
	if s == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return i
}

// GetBool returns the environment variable parsed as bool, or the provided default.
func (c *Config) GetBool(key string, defaultValue ...bool) bool {
	s := c.Get(key)
	if s == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return false
	}
	b, err := strconv.ParseBool(s)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return false
	}
	return b
}