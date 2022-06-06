package common

import "github.com/google/uuid"

func IsNullOrEmpty(value string) bool {
	return len(value) <= 0
}
func GenerateID() string {
	return uuid.New().String()
}
