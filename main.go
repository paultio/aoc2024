package main

import (
	historian "aoc/1"
	nuclear "aoc/2"
	mul "aoc/3"
	wordsearch "aoc/4"
	pageordering "aoc/5"
	"fmt"
	"os"
)

func main() {
	fifth()
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
	unsafe := [][]int{}
	for _, level := range levels {
		if nuclear.CheckLevelSafety(level) {
			safe++
		} else {
			unsafe = append(unsafe, level)
		}
	}
	fmt.Println("Safes: ", safe, " Unsafes: ", len(unsafe))

	// Problem dampener
	dampenedSafe := 0
	for _, level := range unsafe {
		if nuclear.ProblemDampener(level) {
			dampenedSafe++
		}
	}
	fmt.Println("Dampened safes: ", dampenedSafe)
	fmt.Println("Total: ", safe+dampenedSafe)
}

func third() {
	filename := "./3/mul.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(mul.RunMultiplication(string(contents)))
	parts := mul.SplitDoDont(string(contents))
	text := ""
	for _, part := range parts {
		// fmt.Println("Len: ", len(part), "\n\nBEGIN:\n", part, "\nEND")
		text += part
	}
	fmt.Println("Total: ", mul.RunMultiplication(text))
}

func fourth() {
	filename := "./4/wordsearch.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	board := wordsearch.ParseBoard(string(contents))
	// fmt.Println("Board: ", board)
	fmt.Println("Found: ", wordsearch.WordSearch(board, "XMAS"))

	fmt.Println("Crosses: ", wordsearch.FindCrosses(board))
}

func fifth() {
	filename := "./5/page_ordering.txt"
	// filename := "./5/test_data.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	rules, orders := pageordering.ParseFromStr(string(contents))

	indexes := pageordering.ApplyRules(rules, orders)
	fmt.Println("Correct indexes: ", indexes)
	correct_pages := []int{}
	sum := 0
	for i := range indexes {
		// fmt.Println("Correct order: ", orders[i], "middle page: ", orders[i][len(orders[i])/2])
		sum += orders[i][len(orders[i])/2]
		correct_pages = append(correct_pages, orders[i][len(orders[i])/2])
	}
	fmt.Println("Sum: ", sum)
	fmt.Println("Middle pages: ", correct_pages)
}
