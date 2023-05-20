package utils

import (
	"github.com/mozillazg/go-pinyin"
)

func GetShortcut(chinese string) (long, short string) {
	s := pinyin.Convert(chinese, nil)
	a1, a2 := "", ""
	i := 0
	for _, c := range chinese {
		if c >= 0x4e00 && c <= 0x9fa5 {
			a1 = a1 + s[i][0]
			a2 = a2 + string(s[i][0][0])
			i++
		} else {
			a1 = a1 + string(c)
			a2 = a2 + string(c)
		}
	}

	return a1, a2
}
