package aoc

import (
	"strings"
)

func WordSearch(board [][]byte, word string) int {
	found := 0

	// Check if the word is empty
	if len(word) == 0 {
		return 0
	}
	// Check if the board is empty
	if len(board) == 0 {
		return 0
	}
	// Check if the board is empty
	if len(board[0]) == 0 {
		return 0
	}
	// Check if the word is longer than the board
	if len(word) > len(board)*len(board[0]) {
		return 0
	}

	found += searchRows(board, word)
	found += searchColumns(board, word)
	found += searchDiagonals(board, word)

	return found
}

func ParseBoard(boardString string) [][]byte {
	var board [][]byte

	lines := strings.Split(boardString, "\n")
	for _, line := range lines {

		row := []byte{}
		for _, char := range line {
			row = append(row, byte(char))
		}

		board = append(board, row)
	}
	return board
}

func searchLine(line string, word string) int {
	if len(word) == 0 {
		return 0
	}
	if len(line) == 0 {
		return 0
	}
	found := 0
	beginAt := 0
	for i := 0; i < len(line); i++ {
		if line[i] == word[0] {
			beginAt = i
			// Check if the word fits in the line
			if i+len(word) > len(line) {
				return found
			}

			// Check if the word is in the line
			for j := 0; j < len(word); j++ {
				if line[i+j] != word[j] {
					break
				}
				if j == len(word)-1 {
					found++
					i = beginAt
				}
			}
		}
	}

	return found
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func searchRows(board [][]byte, word string) int {
	found := 0
	for row := range board {
		found += searchLine(string(board[row]), word)
		found += searchLine(reverseString(string(board[row])), word)
	}
	return found
}

func searchColumns(board [][]byte, word string) int {
	found := 0
	for col := range board[0] {
		var column []byte
		for row := range board {
			column = append(column, board[row][col])
		}
		found += searchLine(string(column), word)
		found += searchLine(reverseString(string(column)), word)
	}
	return found
}

func searchDiagonals(board [][]byte, word string) int {
	found := 0
	// Top left to bottom right
	for i := 0; i < len(board); i++ {
		var diagonal []byte
		for j := 0; j < len(board)-i; j++ {
			diagonal = append(diagonal, board[j][j+i])
		}
		found += searchLine(string(diagonal), word)
		found += searchLine(reverseString(string(diagonal)), word)
	}
	for i := 1; i < len(board); i++ {
		var diagonal []byte
		for j := 0; j < len(board)-i; j++ {
			diagonal = append(diagonal, board[j+i][j])
		}
		found += searchLine(string(diagonal), word)
		found += searchLine(reverseString(string(diagonal)), word)
	}
	// Top right to bottom left
	for i := 0; i < len(board); i++ {
		var diagonal []byte
		for j := 0; j < len(board)-i; j++ {
			diagonal = append(diagonal, board[j][len(board)-1-j-i])
		}
		found += searchLine(string(diagonal), word)
		found += searchLine(reverseString(string(diagonal)), word)
	}
	for i := 1; i < len(board); i++ {
		var diagonal []byte
		for j := 0; j < len(board)-i; j++ {
			diagonal = append(diagonal, board[j+i][len(board)-1-j])
		}
		found += searchLine(string(diagonal), word)
		found += searchLine(reverseString(string(diagonal)), word)
	}
	return found
}
