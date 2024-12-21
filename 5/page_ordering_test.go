package aoc

import (
	"testing"
)

func TestParseRule(t *testing.T) {
	tests := []struct {
		input    string
		expected Rule
	}{
		{"1|2", Rule{1, 2}},
		{"3|4", Rule{3, 4}},
		{"5|6", Rule{5, 6}},
	}

	for _, test := range tests {
		result := ParseRule(test.input)
		if result != test.expected {
			t.Errorf("ParseRule(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
func TestParseFromStr(t *testing.T) {
	tests := []struct {
		input          string
		expectedRules  []Rule
		expectedOrders [][]int
	}{
		{
			"1|2\n3|4\n\n1,2,3\n4,5,6",
			[]Rule{{1, 2}, {3, 4}},
			[][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			"5|6\n7|8\n\n9,10,11\n12,13,14",
			[]Rule{{5, 6}, {7, 8}},
			[][]int{{9, 10, 11}, {12, 13, 14}},
		},
	}

	for _, test := range tests {
		rules, orders := ParseFromStr(test.input)
		if !equalRules(rules, test.expectedRules) {
			t.Errorf("ParseFromStr(%q) rules = %v; want %v", test.input, rules, test.expectedRules)
		}
		if !equalOrders(orders, test.expectedOrders) {
			t.Errorf("ParseFromStr(%q) orders = %v; want %v", test.input, orders, test.expectedOrders)
		}
	}
}

func TestApply(t *testing.T) {
	tests := []struct {
		rule     Rule
		pages    []int
		expected bool
	}{
		{Rule{1, 2}, []int{1, 2, 3}, true},
		{Rule{2, 1}, []int{1, 2, 3}, false},
		{Rule{1, 3}, []int{1, 2, 3}, true},
		{Rule{3, 1}, []int{1, 2, 3}, false},
		{Rule{1, 2}, []int{3, 2, 1}, false},
		{Rule{2, 3}, []int{3, 2, 1}, false},
		{Rule{1, 4}, []int{1, 2, 3}, false}, // 4 is not in the list
	}

	for _, test := range tests {
		result := test.rule.Apply(test.pages)
		if result != test.expected {
			t.Errorf("Rule(%v).Apply(%v) = %v; want %v", test.rule, test.pages, result, test.expected)
		}
	}
}

func TestApplyRules(t *testing.T) {
	tests := []struct {
		rules    []Rule
		pages    [][]int
		expected []int
	}{
		{
			[]Rule{{1, 2}, {3, 4}},
			[][]int{{1, 2, 3, 4}, {4, 3, 2, 1}, {1, 3, 2, 4}},
			[]int{0, 2},
		},
		{
			[]Rule{{5, 6}, {7, 8}},
			[][]int{{5, 6, 7, 8}, {8, 7, 6, 5}, {5, 7, 6, 8}},
			[]int{0, 2},
		},
		{
			[]Rule{{1, 2}},
			[][]int{{2, 1}, {1, 2}},
			[]int{1},
		},
		{
			[]Rule{{1, 2}, {2, 3}},
			[][]int{{1, 2, 3}, {3, 2, 1}, {1, 3, 2}},
			[]int{0},
		},
	}

	for _, test := range tests {
		result := ApplyRules(test.rules, test.pages)
		if !equalSlices(result, test.expected) {
			t.Errorf("ApplyRules(%v, %v) = %v; want %v", test.rules, test.pages, result, test.expected)
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

func equalRules(a, b []Rule) bool {
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

func equalOrders(a, b [][]int) bool {
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
