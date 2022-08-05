package crypto

import (
	"github.com/google/uuid"
)

func GetUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
