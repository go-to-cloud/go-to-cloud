package utils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// TestComSet 补集，无交集
func TestComSet1(t *testing.T) {
	oldSet := New[int](3, 4, 5)
	newSet := New[int](1, 0)

	sub := Minus(oldSet, newSet) // expected: 4, 5
	assert.True(t, reflect.DeepEqual(SortList(sub), []int{3, 4, 5}))

	com := Complement(oldSet, newSet)
	assert.True(t, reflect.DeepEqual(SortList(com), []int{0, 1}))
}

// TestComSet 补集，有交集
func TestComSet2(t *testing.T) {
	oldSet := New[int](3, 4, 5)
	newSet := New[int](1, 3)

	sub := Minus(oldSet, newSet) // expected: 4, 5
	assert.True(t, reflect.DeepEqual(SortList(sub), []int{4, 5}))

	com := Complement(oldSet, newSet)
	assert.True(t, reflect.DeepEqual(SortList(com), []int{1}))
}

// TestSets 集合测试
func TestSets(t *testing.T) {
	s := New[int](1, 2, 3)

	if Count(s) != 3 {
		t.Errorf("Expected 3, got %d", Count(s))
	}

	if !Has(s, 1, 2, 3) {
		t.Errorf("Expected true, got false")
	}

	if Has(s, 4) {
		t.Errorf("Expected false, got true")
	}

	Remove(s, 2)
	if Count(s) != 2 {
		t.Errorf("Expected 2, got %d", Count(s))
	}

	if !reflect.DeepEqual(SortList(s), []int{1, 3}) {
		t.Errorf("Expected [1, 3], got %v", List(s))
	}

	s2 := New[int](3, 4, 5)
	u := Union(s, s2)
	SortList(u)
	su := New[int](1, 3, 4, 5)
	SortList(su)
	if !reflect.DeepEqual(u, su) {
		t.Errorf("Expected [1, 3, 4, 5], got %v", List(u))
	}

	m := Minus(s, s2)
	if !reflect.DeepEqual(List(m), []int{1}) {
		t.Errorf("Expected [1], got %v", List(m))
	}

	c := Complement(s, New[int](1, 2, 3, 4, 5))
	if !reflect.DeepEqual(SortList(c), []int{2, 4, 5}) {
		t.Errorf("Expected [2, 4, 5], got %v", List(c))
	}
}

func TestIntersectGeneric(t *testing.T) {
	// Test with empty sets
	s1 := New[int]()
	s2 := New[int]()
	s3 := New[int]()
	expected := New[int]()

	assert.True(t, Empty(s1))

	actual := IntersectGeneric(s1, s2, s3)
	assert.True(t, reflect.DeepEqual(actual, expected))

	// Test with non-empty sets
	xs1 := New[string]("a", "b", "c")
	xs2 := New[string]("a", "c", "d")
	xs3 := New[string]("a", "b", "d")
	xexpected := New[string]("a")

	xactual := IntersectGeneric[string](xs1, xs2, xs3)
	assert.True(t, reflect.DeepEqual(xactual, xexpected))

	Clear(xs3)
	assert.True(t, Empty(s3))
}
