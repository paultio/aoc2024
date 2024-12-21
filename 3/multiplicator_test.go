package aoc

import (
	"testing"
)

func TestFindMulExpressions(t *testing.T) {
	tests := []struct {
		text     string
		expected []string
	}{
		{"This is a test with mul(3,4) expression", []string{"mul(3,4)"}},
		{"No multiplication here", []string{}},
		{"Multiple mul(1,2) and mul(5,6) expressions", []string{"mul(1,2)", "mul(5,6)"}},
		{"Nested mul(7,8)mul(9,10) expressions", []string{"mul(7,8)", "mul(9,10)"}},
		{"Invalid mul(11,12,13) expression", []string{}},
		{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}},
	}

	for _, test := range tests {
		result := findMulExpressions(test.text)
		if len(result) != len(test.expected) {
			t.Errorf("findMulExpressions(%v) = %v; expected %v", test.text, result, test.expected)
		}
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("findMulExpressions(%v) = %v; expected %v", test.text, result, test.expected)
			}
		}
		// if !reflect.DeepEqual(result, test.expected) {
		// 	t.Errorf("findMulExpressions(%v) = %v; expected %v", test.text, result, test.expected)
		// }
	}
}

func TestExtractNumbers(t *testing.T) {
	tests := []struct {
		expression string
		expected   []int
	}{
		{"mul(3,4)", []int{3, 4}},
		{"mul(10,20)", []int{10, 20}},
		{"mul(0,0)", []int{0, 0}},
		{"mul(123,456)", []int{123, 456}},
		{"mul(7,8)", []int{7, 8}},
	}

	for _, test := range tests {
		result := extractNumbers(test.expression)
		if len(result) != len(test.expected) {
			t.Errorf("extractNumbers(%v) = %v; expected %v", test.expression, result, test.expected)
		}
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("extractNumbers(%v) = %v; expected %v", test.expression, result, test.expected)
			}
		}
	}
}
func TestRunMultiplication(t *testing.T) {
	tests := []struct {
		text     string
		expected int
	}{
		{"This is a test with mul(3,4) expression", 12},
		{"No multiplication here", 0},
		{"Multiple mul(1,2) and mul(5,6) expressions", 32},
		{"Nested mul(7,8)mul(9,10) expressions", 146},
		{"Invalid mul(11,12,13) expression", 0},
		{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", 161},
	}

	for _, test := range tests {
		result := RunMultiplication(test.text)
		if result != test.expected {
			t.Errorf("RunMultiplication(%v) = %v; expected %v", test.text, result, test.expected)
		}
	}
}
