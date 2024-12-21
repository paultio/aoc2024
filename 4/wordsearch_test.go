package aoc

import "testing"

func TestSearchLine(t *testing.T) {
	tests := []struct {
		line     string
		word     string
		expected int
	}{
		{"hello", "he", 1},
		{"hello", "lo", 1},
		{"hello", "hello", 1},
		{"hello", "world", 0},
		{"hellohello", "hello", 2},
		{"", "hello", 0},
		{"hello", "", 0},
		{"aaaaa", "aa", 4},
	}

	for _, test := range tests {
		result := searchLine(test.line, test.word)
		if result != test.expected {
			t.Errorf("searchLine(%q, %q) = %d; expected %d", test.line, test.word, result, test.expected)
		}
	}
}
func TestSearchRows(t *testing.T) {
	tests := []struct {
		board    [][]byte
		word     string
		expected int
	}{
		{[][]byte{{}, {'h', 'e', 'l', 'l', 'o'}}, "he", 1},
		{[][]byte{{'l', 'o'}, {'h', 'e', 'l', 'l', 'o'}}, "lo", 2},
		{[][]byte{{'h', 'e', 'l', 'l', 'o'}}, "hello", 1},
		{[][]byte{{'h', 'e', 'l', 'l', 'o'}}, "ll", 2},
		{[][]byte{{'h', 'e', 'l', 'l', 'o'}, {'h', 'e', 'l', 'l', 'o'}}, "hello", 2},
		{[][]byte{{'h', 'e', 'l', 'l', 'o'}, {'o', 'l', 'l', 'e', 'h'}}, "hello", 2},
		{[][]byte{{'a', 'a', 'a', 'a', 'a'}}, "aa", 8},
		{[][]byte{}, "hello", 0},
		{[][]byte{{'h', 'e', 'l', 'l', 'o'}}, "", 0},
	}

	for _, test := range tests {
		result := searchRows(test.board, test.word)
		if result != test.expected {
			t.Errorf("searchRows(%q, %q) = %d; expected %d", test.board, test.word, result, test.expected)
		}
	}
}
func TestSearchColumns(t *testing.T) {
	tests := []struct {
		board    [][]byte
		word     string
		expected int
	}{
		{[][]byte{{'h', 'e', 'l', 'l', 'o'}, {'h', 'e', 'l', 'l', 'o'}}, "he", 0},
		{[][]byte{{'h', 'e', 'l', 'l', 'o'}, {'o', 'l', 'l', 'e', 'h'}}, "ho", 2},
		{[][]byte{{'h', 'e', 'l', 'l', 'o'}, {'o', 'l', 'l', 'e', 'h'}}, "eh", 0},
		{[][]byte{{'h', 'e', 'l', 'l', 'o'}, {'o', 'l', 'l', 'e', 'h'}}, "ll", 2},
		{[][]byte{{'h', 'e', 'l', 'l', 'o'}, {'h', 'e', 'l', 'l', 'o'}}, "ll", 4},
		{[][]byte{{'a', 'b', 'a', 'd', 'e'}, {'b', 'b', 'b', 'd', 'e'}, {'c', 'b', 'c', 'd', 'e'}}, "abc", 2},
		{[][]byte{{'a', 'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a', 'a'}}, "aa", 20},
		{[][]byte{{'h', 'e', 'l', 'l', 'o'}}, "", 0},
	}

	for _, test := range tests {
		result := searchColumns(test.board, test.word)
		if result != test.expected {
			t.Errorf("searchColumns(%q, %q) = %d; expected %d", test.board, test.word, result, test.expected)
		}
	}
}
func TestSearchDiagonals(t *testing.T) {
	tests := []struct {
		board    [][]byte
		word     string
		expected int
	}{
		{[][]byte{{'h', 'e', 'l', 'l', 'o'}, {'e', 'l', 'l', 'o', 'h'}, {'l', 'l', 'o', 'h', 'e'}, {'l', 'o', 'h', 'e', 'l'}, {'o', 'h', 'e', 'l', 'l'}}, "hloel", 1},
		{[][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}}, "aei", 1},
		{[][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}}, "xyz", 0},
		{[][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}}, "", 0},
		{[][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'c', 'b', 'a'}}, "aea", 2},
	}

	for _, test := range tests {
		result := searchDiagonals(test.board, test.word)
		if result != test.expected {
			t.Errorf("searchDiagonals(%q, %q) = %d; expected %d", test.board, test.word, result, test.expected)
		}
	}
}
