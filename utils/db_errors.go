package utils

import (
	"errors"

	"gorm.io/gorm"
)

func HandleDBError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("Record Not Found")
	}
	return nil
}
