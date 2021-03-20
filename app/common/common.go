package common

import (
	"crypto/rand"
	"fmt"
	"os"
)

// GererateUUID generate and return a UUID.
func GenerateUUID() (uuid string) {
	b := make([]byte, 16)

	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Unable to create UUID, error: ", err)
		return
	}

	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid
}

// GetEnv return a environment variable passing a default value.
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
