package config
import (
	"os"
	"github.com/joho/godotenv"
)
type Config struct{
	AppPort   string
    DBHost    string
    DBUser    string
    DBPass    string
    DBName    string
    DBPort    string
    JWTSecret string
}
func Load() *Config{
	godotenv.Load()

	return &Config{
		 AppPort:  os.Getenv("APP_PORT"),
        DBHost:    os.Getenv("DB_HOST"),
        DBUser:    os.Getenv("DB_USER"),
        DBPass:    os.Getenv("DB_PASS"),
        DBName:    os.Getenv("DB_NAME"),
        DBPort:    os.Getenv("DB_PORT"),
        JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
