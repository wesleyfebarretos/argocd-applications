package config

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type DatabaseConfig struct {
	ConnURI  string
	MaxConns int
}

type Config struct {
	Database DatabaseConfig
}

var (
	config *Config
	once   sync.Once
)

func getEnvOrThrow(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

func getEnvAsInt(key string, dft int) int {
	value, ok := os.LookupEnv(key)

	if !ok {
		return dft
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Environment variable %s is not a valid integer", key)
	}

	return result
}

func Init() {
	once.Do(func() {
		config = &Config{
			Database: DatabaseConfig{
				ConnURI:  getEnvOrThrow("DB_URI"),
				MaxConns: getEnvAsInt("DB_MAX_CONNS", 10),
			},
		}
	})
}

func Get() *Config {
	if config == nil {
		panic("config.Init() must be called before config.Get()")
	}

	return config
}
