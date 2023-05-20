package utils

import (
	"math/rand"
)

const asciiOffset = 33
const asciiMax = 126 - asciiOffset // 用于密码的ascii位于33～126之间

// StrongPasswordGen 生成强密码，len表示密码长度
func StrongPasswordGen(len uint8) *string {
	if len < 6 {
		len = 6
	}

	pwd := make([]byte, len)
	for i := 0; i < int(len); i++ {
		pwd[i] = byte(rand.Intn(asciiMax+1) + asciiOffset)
	}

	t := string(pwd)
	return &t
}
