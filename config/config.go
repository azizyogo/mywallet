package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string
	GinMode    string

	MySQLDSN          string
	MySQLMaxIdleConns int
	MySQLMaxOpenConns int

	JWTSecret          string
	JWTExpirationHours int
}

func LoadConfig() Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Try to read .env file, but don't fail if it doesn't exist
	// In Docker, we use environment variables passed by docker-compose
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found, using environment variables: %v", err)
	}

	// Set defaults
	viper.SetDefault("JWT_EXPIRATION_HOURS", 24)
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("GIN_MODE", "debug")

	return Config{
		ServerPort: viper.GetString("SERVER_PORT"),
		GinMode:    viper.GetString("GIN_MODE"),

		MySQLDSN:          viper.GetString("MYSQL_DSN"),
		MySQLMaxIdleConns: viper.GetInt("MYSQL_MAX_IDLE_CONNS"),
		MySQLMaxOpenConns: viper.GetInt("MYSQL_MAX_OPEN_CONNS"),

		JWTSecret:          viper.GetString("JWT_SECRET"),
		JWTExpirationHours: viper.GetInt("JWT_EXPIRATION_HOURS"),
	}
}
