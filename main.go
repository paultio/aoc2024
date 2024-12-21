package main

import (
	historian "aoc/1"
	nuclear "aoc/2"
	"fmt"
)

func main() {
	second()
}

func first() {
	l1, l2, err := historian.ReadIntsFromFile("./1/lists.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Lists distance: ", historian.ListCompare(l1, l2))

	fmt.Println("Similarity: ", historian.SimilarityScore(l1, l2))
}

func second() {
	levels, err := nuclear.ReadLevelsFromFile("./2/levels.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	safe := 0
	for _, level := range levels {
		if nuclear.CheckLevelSafety(level) {
			safe++
		}
	}
	fmt.Println("Safes: ", safe)
}
