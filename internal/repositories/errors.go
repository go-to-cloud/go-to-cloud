package repositories

import "gorm.io/gorm"
import "errors"

func isNoRecordError(err error) bool {
	return err != nil && errors.Is(err, gorm.ErrRecordNotFound)
}

func returnWithError[T interface{}](ret T, err error) (T, error) {
	if err == nil {
		return ret, nil
	}

	if isNoRecordError(err) {
		return ret, nil
	} else {
		return ret, err
	}
}
