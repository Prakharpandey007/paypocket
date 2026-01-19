package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort   string
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBPort    string
	JWTSecret string
}

func Load() *Config {
	// Try to load .env file from project root
	// Find project root by looking for go.mod file
	envPath := findEnvFile()
	
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Warning: Error loading .env file from %s: %v. Using environment variables if set.", envPath, err)
	}

	cfg := &Config{
		AppPort:   getEnv("APP_PORT", "8080"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBUser:    getEnv("DB_USER", "root"),
		DBPass:    getEnv("DB_PASS", ""),
		DBName:    getEnv("DB_NAME", "paypocket"),
		DBPort:    getEnv("DB_PORT", "3306"),
		JWTSecret: getEnv("JWT_SECRET", "supersecretkey"),
	}

	// Validate critical config values
	if cfg.DBPort == "" {
		log.Fatal("DB_PORT is required but not set")
	}
	if cfg.DBHost == "" {
		log.Fatal("DB_HOST is required but not set")
	}

	return cfg
}

func findEnvFile() string {
	// Start from current directory and walk up to find go.mod
	// This ensures we find the project root regardless of where the binary is run from
	wd, err := os.Getwd()
	if err != nil {
		return ".env"
	}

	// Check current directory and parent directories for .env file
	dir := wd
	for i := 0; i < 10; i++ { // Limit to 10 levels up
		envPath := dir + "/.env"
		if _, err := os.Stat(envPath); err == nil {
			return envPath
		}
		// Check if we've reached the project root (has go.mod)
		if _, err := os.Stat(dir + "/go.mod"); err == nil {
			// Found project root, .env should be here
			return dir + "/.env"
		}
		// Move up one directory
		parent := dir + "/.."
		parentAbs, err := filepath.Abs(parent)
		if err != nil || parentAbs == dir {
			break // Can't go up further
		}
		dir = parentAbs
	}

	// Fallback to current directory
	return ".env"
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
