package env

import (
	"os"
	"strconv"
)

func GetString(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GetInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		i, err := strconv.Atoi(value)
		if err == nil {
			return i
		}
	}
	return defaultValue
}