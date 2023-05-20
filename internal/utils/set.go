package utils

import (
	"sync"
)

type TSet interface {
	int | uint | int64 | uint64 | float64 | float32 | string
}

type CollectionMap[T TSet] map[T]bool

type Set[T TSet] struct {
	sync.RWMutex
	m CollectionMap[T]
}

// New 新建集合对象
func New[T TSet](items ...T) *Set[T] {
	s := &Set[T]{
		m: make(CollectionMap[T], len(items)),
	}
	Add(s, items...)
	return s
}

// Add 添加元素
func Add[T TSet](s *Set[T], items ...T) {
	s.Lock()
	defer s.Unlock()
	for _, v := range items {
		s.m[v] = true
	}
}

// Remove 删除元素
func Remove[T TSet](s *Set[T], items ...T) {
	s.Lock()
	defer s.Unlock()
	for _, v := range items {
		delete(s.m, v)
	}
}

// Has 判断元素是否存在
func Has[T TSet](s *Set[T], items ...T) bool {
	s.RLock()
	defer s.RUnlock()
	for _, v := range items {
		if _, ok := s.m[v]; !ok {
			return false
		}
	}
	return true
}

// Count 元素个数
func Count[T TSet](s *Set[T]) int {
	return len(s.m)
}

// Clear 清空集合
func Clear[T TSet](s *Set[T]) {
	s.Lock()
	defer s.Unlock()
	s.m = map[T]bool{}
}

// Empty 空集合判断
func Empty[T TSet](s *Set[T]) bool {
	return len(s.m) == 0
}

// List 无序列表
func List[T TSet](s *Set[T]) []T {
	s.RLock()
	defer s.RUnlock()
	list := make([]T, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func bubbleSort[T TSet](x []T) {
	n := len(x)
	for {
		swapped := false
		for i := 1; i < n; i++ {
			if x[i] < x[i-1] {
				x[i-1], x[i] = x[i], x[i-1]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

// SortList 排序列表
func SortList[T TSet](s *Set[T]) []T {
	s.RLock()
	defer s.RUnlock()
	list := make([]T, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	bubbleSort(list)
	return list
}

// Union 并集
func Union[T TSet](s *Set[T], sets ...*Set[T]) *Set[T] {
	r := New(List(s)...)
	for _, set := range sets {
		for e := range set.m {
			r.m[e] = true
		}
	}
	return r
}

// Minus 差集
func Minus[T TSet](s *Set[T], sets ...*Set[T]) *Set[T] {
	r := New(List(s)...)
	for _, set := range sets {
		for e := range set.m {
			if _, ok := s.m[e]; ok {
				delete(r.m, e)
			}
		}
	}
	return r
}

// IntersectGeneric 泛型交集
func IntersectGeneric[T TSet](s *Set[T], sets ...*Set[T]) *Set[T] {
	r := New(List(s)...)
	for _, set := range sets {
		for e := range s.m {
			if _, ok := set.m[e]; !ok {
				delete(r.m, e)
			}
		}
	}
	return r
}

// Complement 补集
func Complement[T TSet](s *Set[T], full *Set[T]) *Set[T] {
	r := New[T]()
	for e := range full.m {
		if _, ok := s.m[e]; !ok {
			Add(r, e)
		}
	}
	return r
}
