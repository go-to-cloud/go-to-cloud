package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrongPasswordGen(t *testing.T) {
	a1 := StrongPasswordGen(3)
	assert.Equal(t, 6, len(*a1))

	a2 := StrongPasswordGen(6)
	assert.Equal(t, 6, len(*a2))

	a3 := StrongPasswordGen(8)
	assert.Equal(t, 8, len(*a3))
}
