package utils

import "github.com/google/uuid"

func genUUID() string {
	return uuid.New().String()
}
