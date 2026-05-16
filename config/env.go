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
	host := getString("DB_HOST", "localhost")
    port := getString("DB_PORT", "5432")
    user := getString("POSTGRES_USER", "postgres")
    password := getString("POSTGRES_PASSWORD", "")
    dbname := getString("POSTGRES_DB", "postgres")
	sslmode := getString("DB_SSLMODE", "disable")
	url := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        host, port, user, password, dbname, sslmode,
    )
	return &DatabaseConfig{
		Url: url,
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