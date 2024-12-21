package aoc

import (
	"regexp"
	"strconv"
	"strings"
)

func findMulExpressions(text string) []string {
	// Regular expression to match `mul(int,int)` exactly
	pattern := `mul\([0-9]+,[0-9]+\)`
	re := regexp.MustCompile(pattern)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)

	return matches
}

func RunMultiplication(text string) int {
	// Find all mul expressions
	expressions := findMulExpressions(text)

	// Evaluate the expressions
	result := 0
	// fmt.Println("Text:", text)
	// fmt.Println("Found expressions:", expressions)
	for _, expr := range expressions {
		// Extract the numbers from the expression
		nums := extractNumbers(expr)

		// Multiply the numbers
		// fmt.Println(expr, "->", nums, "->", result, "+", nums[0], "*", nums[1], "=", result+nums[0]*nums[1])
		result += nums[0] * nums[1]
	}

	return result
}

func SplitDoDont(text string) []string {
	// Split the text by "do_not_mul"
	parts := strings.Split(text, "do()")
	ok := []string{}
	for _, part := range parts {
		moreParts := strings.Split(part, "don't()")
		ok = append(ok, moreParts[0])
	}
	return ok
}

func extractNumbers(mul string) []int {
	strNums := strings.Split(strings.Split(strings.Split(mul, "(")[1], ")")[0], ",")
	nums := make([]int, 2)
	for i, strNum := range strNums {
		num, _ := strconv.Atoi(strNum)
		nums[i] = num
	}

	return nums
}
