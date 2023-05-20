package utils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestDistinct(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected interface{}
	}{
		{[]int{1, 2, 2, 3, 3, 3}, []int{1, 2, 3}},
		{[]string{"apple", "banana", "banana", "cherry"}, []string{"apple", "banana", "cherry"}},
		{[]uint{1, 2, 3, 3, 2, 1}, []uint{1, 2, 3}},
	}

	result1 := Distinct(testCases[0].input.([]int))
	assert.True(t, reflect.DeepEqual(result1, testCases[0].expected))

	result2 := Distinct(testCases[1].input.([]string))
	assert.True(t, reflect.DeepEqual(result2, testCases[1].expected))

	result3 := Distinct(testCases[2].input.([]uint))
	assert.True(t, reflect.DeepEqual(result3, testCases[2].expected))
}
