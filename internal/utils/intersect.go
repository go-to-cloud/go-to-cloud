package utils

func Intersect(a, b []uint) []uint {
	counter := make(map[uint]bool)
	rlt := make([]uint, 0)
	for _, a := range a {
		counter[a] = true
	}
	for _, b := range b {
		sz, _ := counter[b]
		if sz {
			rlt = append(rlt, b)
		}
	}
	return rlt
}
