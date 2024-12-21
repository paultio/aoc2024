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

func TestFindLocations(t *testing.T) {
	tests := []struct {
		line     string
		word     string
		expected [][2]int
	}{
		{"hello", "he", [][2]int{{0, 1}}},
		{"hello", "lo", [][2]int{{3, 4}}},
		{"hello", "hello", [][2]int{{0, 4}}},
		{"hello", "world", [][2]int{}},
		{"hellohello", "hello", [][2]int{{0, 4}, {5, 9}}},
		{"", "hello", [][2]int{}},
		{"hello", "", [][2]int{}},
		{"aaaaa", "aa", [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}},
	}

	for _, test := range tests {
		result := findLocations(test.line, test.word)
		if !equalLocations(result, test.expected) {
			t.Errorf("findLocations(%q, %q) = %v; expected %v", test.line, test.word, result, test.expected)
		}
	}
}

func TestFindCrosses(t *testing.T) {
	tests := []struct {
		board    [][]byte
		expected int
	}{
		{[][]byte{
			{'S', 'S', 'S'},
			{'S', 'A', 'S'},
			{'S', 'S', 'S'},
		}, 0},
		{[][]byte{
			{'M', 'S', 'M'},
			{'S', 'A', 'S'},
			{'M', 'S', 'M'},
		}, 0},
		{[][]byte{
			{'M', 'S', 'S'},
			{'S', 'A', 'S'},
			{'M', 'S', 'S'},
		}, 1},
		{[][]byte{
			{'M', 'S', 'M', 'S'},
			{'S', 'A', 'S', 'A'},
			{'S', 'S', 'S', 'S'},
			{'S', 'A', 'S', 'A'},
		}, 1},

		{[][]byte{
			{'M', 'S', 'M', 'S', 'M'},
			{'S', 'A', 'S', 'A', 'S'},
			{'S', 'S', 'S', 'S', 'S'},
			{'S', 'A', 'S', 'A', 'S'},
			{'M', 'S', 'M', 'S', 'M'},
		}, 4},
	}

	for _, test := range tests {
		result := FindCrosses(test.board)
		if result != test.expected {
			t.Errorf("findCrosses(%q) = %d; expected %d", test.board, result, test.expected)
		}
	}
}

func equalLocations(a, b [][2]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i][0] != b[i][0] || a[i][1] != b[i][1] {
			return false
		}
	}
	return true
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) || a[i][0] != b[i][0] || a[i][1] != b[i][1] {
			return false
		}
	}
	return true
}
