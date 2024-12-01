package memoize

import "fmt"

type memoized struct {
	f     func(int, []int) int
	cache map[int]int
}

func (m *memoized) Call(x int, list []int) int {
	if v, ok := m.cache[x]; ok {
		fmt.Println("cache hit")
		return v
	}
	result := m.f(x, list)
	m.cache[x] = result
	return result
}

func Memoize(f func(int, []int) int) *memoized {
	return &memoized{f: f, cache: make(map[int]int)}
}
