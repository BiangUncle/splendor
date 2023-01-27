package utils

import uuid "github.com/satori/go.uuid"

func GetUuidV1() string {
	return uuid.NewV1().String()
}

func GetUuidV4() string {
	return uuid.NewV4().String()
}
