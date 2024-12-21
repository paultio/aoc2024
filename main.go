package main

import (
	aoc "aoc/1"
	"fmt"
)

func main() {
	first()
}

func first() {
	l1, l2, err := aoc.ReadIntsFromFile("./1/lists.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Lists distance: ", aoc.ListCompare(l1, l2))

	fmt.Println("Similarity: ", aoc.SimilarityScore(l1, l2))
}
