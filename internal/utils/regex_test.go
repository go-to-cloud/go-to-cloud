package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	validEmails := []string{
		"test@example.com",
		"test.user@example.com",
		"test.user+123@example.com",
	}
	invalidEmails := []string{
		"test@example",
		"test@example.",
		"test@.com",
		"test@com.",
		"test@com",
	}
	for _, email := range validEmails {
		if !IsValidEmail(email) {
			t.Errorf("IsValidEmail(%q) = false, expected true", email)
		}
	}
	for _, email := range invalidEmails {
		if IsValidEmail(email) {
			t.Errorf("IsValidEmail(%q) = true, expected false", email)
		}
	}
}

func TestIsValidMobile(t *testing.T) {
	validMobiles := []string{
		"13812345678",
		"13987654321",
	}
	invalidMobiles := []string{
		"123456789",
		"138123456789",
		"1381234567a",
	}
	for _, mobile := range validMobiles {
		assert.True(t, IsValidMobile(mobile))
	}
	for _, mobile := range invalidMobiles {
		assert.False(t, IsValidMobile(mobile))
	}
}
