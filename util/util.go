package util

import (
	"time"
)

func InitExpensiveFuncWithCache() func(p int) int {
	var cache = make(map[int]int)
	return func(p int) int {
		if ret, ok := cache[p]; ok {
			return ret
		}
		// expensive computation
		time.Sleep(1 * time.Second)
		r := p * 2
		cache[p] = r
		return r
	}
}
