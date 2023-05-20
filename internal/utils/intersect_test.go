package utils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	testCases := []struct {
		a, b, want []uint
	}{
		{[]uint{1, 2, 3}, []uint{2, 3, 4}, []uint{2, 3}},
		{[]uint{1, 2, 3}, []uint{4, 5, 6}, []uint{}},
		{[]uint{1, 2, 3}, []uint{}, []uint{}},
		{[]uint{}, []uint{2, 3, 4}, []uint{}},
		{[]uint{}, []uint{}, []uint{}},
	}
	for _, tc := range testCases {
		got := Intersect(tc.a, tc.b)
		assert.True(t, reflect.DeepEqual(got, tc.want))
	}
}
