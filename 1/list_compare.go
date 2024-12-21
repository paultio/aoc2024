package aoc

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func ReadIntsFromFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var l1, l2 []int
	for {
		var num1, num2 int
		_, err := fmt.Fscanf(file, "%d    %d\n", &num1, &num2)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, nil, err
		}
		l1 = append(l1, num1)
		l2 = append(l2, num2)
	}
	return l1, l2, nil
}

func ListCompare(l1 []int, l2 []int) int {
	totalDiff := 0
	sort.Ints(l1)
	sort.Ints(l2)
	for idx, num1 := range l1 {
		num2 := l2[idx]
		diff := num1 - num2
		if diff < 0 {
			diff = -diff
		}
		totalDiff += diff
	}
	return totalDiff
}

func SimilarityScore(l1 []int, l2 []int) int {
	similarity := 0

	for _, num := range l1 {
		similarity += num * nTimesInList(num, l2)
	}

	return similarity
}

func nTimesInList(n int, list []int) int {
	var cache = make(map[int]int)
	if ret, ok := cache[n]; ok {
		return ret
	}

	for _, num := range list {
		if num == n {
			cache[n]++
		}
	}
	return cache[n]
}
