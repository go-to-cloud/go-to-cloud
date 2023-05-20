package utils

import (
	"regexp"
	"strings"
)

func isMatch(pattern, o string) bool {
	match, err := regexp.MatchString(pattern, o)
	if err != nil {
		return false
	}
	return match
}

func IsValidEmail(email string) bool {
	return isMatch(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, strings.ToLower(email))
}

func IsValidMobile(mobile string) bool {
	return isMatch(`^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`, mobile)
}
