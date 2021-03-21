package common

import (
	"crypto/rand"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
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

// IsEmailValid check if email there is valid body
func IsEmailValid(email string) bool {
	pattern := "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	regex := regexp.MustCompile(pattern)

	if len(email) < 3 || len(email) > 254 {
		return false
	}

	// Check email string with regex
	if !regex.MatchString(email) {
		return false
	}

	// Check domain
	domain := strings.Split(email, "@")[1]
	mx, err := net.LookupMX(domain)
	if err != nil || len(mx) == 0 {
		return false
	}

	return true
}
