package aoc

import (
	"testing"
)

func TestListCompare(t *testing.T) {
	tests := []struct {
		l1       []int
		l2       []int
		expected int
	}{
		{[]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}, 11},
		{[]int{1, 3, 3}, []int{1, 2, 3}, 1},
	}

	for _, test := range tests {
		result := ListCompare(test.l1, test.l2)
		if result != test.expected {
			t.Errorf("listCompare(%v, %v) = %v; expected %v", test.l1, test.l2, result, test.expected)
		}
	}
}
func TestReadIntsFromFile(t *testing.T) {
	tests := []struct {
		filename string
		l1       []int
		l2       []int
		err      error
	}{
		{"./testlists.txt", []int{6, 2, 7}, []int{9, 8, 7}, nil},
	}

	for _, test := range tests {
		l1, l2, err := ReadIntsFromFile(test.filename)
		if err != nil && test.err == nil {
			t.Errorf("ReadIntsFromFile(%v) returned unexpected error: %v", test.filename, err)
		}
		if err == nil && test.err != nil {
			t.Errorf("ReadIntsFromFile(%v) expected error: %v, got nil", test.filename, test.err)
		}
		if !equalSlices(l1, test.l1) || !equalSlices(l2, test.l2) {
			t.Errorf("ReadIntsFromFile(%v) = (%v, %v); expected (%v, %v)", test.filename, l1, l2, test.l1, test.l2)
		}
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
