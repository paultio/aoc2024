package aoc

import (
	"testing"
)

func TestCheckLevelSafety(t *testing.T) {
	tests := []struct {
		levels   []int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, true},
		{[]int{4, 3, 2, 1}, true},
		{[]int{1, 3, 2, 4}, false},
		{[]int{1, 2, 5}, true},
		{[]int{1, 2, 3, 6}, true},
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, true},
		{[]int{1, 2, 3, 2, 1}, false},
		// From aoc
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, test := range tests {
		result := CheckLevelSafety(test.levels)
		if result != test.expected {
			t.Errorf("checkLevelSafety(%v) = %v; expected %v", test.levels, result, test.expected)
		}
	}
}

func TestProblemDampener(t *testing.T) {
	tests := []struct {
		levels   []int
		expected bool
	}{
		// From aoc
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{8, 6, 4, 4, 1}, true},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, test := range tests {
		result := ProblemDampener(test.levels)
		if result != test.expected {
			t.Errorf("checkLevelSafety(%v) = %v; expected %v", test.levels, result, test.expected)
		}
	}
}

func TestReadLevelsFromFile(t *testing.T) {
	tests := []struct {
		filename string
		expected [][]int
		hasError bool
	}{
		{"./testlevels.txt", [][]int{{1, 2, 3}, {1, 2}, {1, 7}}, false},
	}

	for _, test := range tests {
		result, err := ReadLevelsFromFile(test.filename)
		if (err != nil) != test.hasError {
			t.Errorf("ReadLevelsFromFile(%v) error = %v; expected error = %v", test.filename, err, test.hasError)
		}
		if !test.hasError && !equal(result, test.expected) {
			t.Errorf("ReadLevelsFromFile(%v) = %v; expected %v", test.filename, result, test.expected)
		}
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
