package env

import (
	"os"
	"strconv"
)

func GetStringEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetIntEnv(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {

		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
