package repositories

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestIsNoRecordError(t *testing.T) {
	err := gorm.ErrRecordNotFound

	var tmp *int
	_, e := returnWithError(tmp, err)

	assert.Nil(t, e)
}
