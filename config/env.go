package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Init() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env: %w", err)
	}
	return nil
}

type DatabaseConfig struct {
	Url string
}

func NewDatabaseConfig() *DatabaseConfig{
	return &DatabaseConfig{
		Url: getString("DB_URL", "localhost"),
	}
}

type LogConfig struct {
	Level int
	Format string
	LevelStr string
	FilePath string
}

func NewLogConfig() *LogConfig{
	return &LogConfig{
		Level: getInt("LOG_LEVEL", 0),
		LevelStr: getString("LOG_LEVEL_STR", "info"),
		Format: getString("LOG_FORMAT", "json"),
		FilePath: getString("LOG_PATH", "logs/app.log"),
	}
}

func DefaultLogConfig() *LogConfig {
    return &LogConfig{
        Level:  0, // Info
        Format: "console",
    }
}

func getString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}

func getBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	b, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return b
}