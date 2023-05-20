package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChinesePinyin(t *testing.T) {
	a := "中国人"
	l, s := GetShortcut(a)
	assert.Equal(t, "zhongguoren", l)
	assert.Equal(t, "zgr", s)

	a = "Chines中aa"
	l, s = GetShortcut(a)
	assert.Equal(t, "Chineszhongaa", l)
	assert.Equal(t, "Chineszaa", s)
}
