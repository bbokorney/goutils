package goutils

import (
	"log"
	"os"
)

// GetRequiredEnv tries to get an environment variable.
// If the variable is not set, it exits the process.
func GetRequiredEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing required ENV: %s", key)
	}
	return val
}
