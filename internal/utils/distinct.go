package utils

func Distinct[T string | int | uint](a []T) []T {
	distinct := make([]T, 0)
	o := make(map[T]bool)

	for _, t := range a {
		if _, ok := o[t]; !ok {
			o[t] = true
			distinct = append(distinct, t)
		}
	}
	return distinct
}
