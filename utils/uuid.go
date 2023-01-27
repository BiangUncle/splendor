package utils

import uuid "github.com/satori/go.uuid"

func GetUuidV1() string {
	return uuid.NewV1().String()
}

func GetUuidV4() string {
	return uuid.NewV4().String()
}

func CompressUuid(id string) string {
	if len(id) <= 4 {
		return id
	}

	return id[:2] + "**" + id[len(id)-2:]
}
