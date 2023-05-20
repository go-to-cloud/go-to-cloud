package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAesEny(t *testing.T) {
	plaintext := "Hello中文"

	encoded := AesEny(plaintext)

	decoded := AesEny(string(encoded))

	assert.Equal(t, plaintext, string(decoded))
}

func TestBase64AesEny(t *testing.T) {
	plaintext := "hello world"
	encoded := Base64AesEny(plaintext)
	decoded := Base64AesEnyDecode(encoded)

	assert.Equal(t, decoded, plaintext)
}
