package aoc

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadLevelsFromFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Split the line by spaces
		parts := strings.Fields(line)
		var intList []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("error converting %q to integer: %v", part, err)
			}
			intList = append(intList, num)
		}
		result = append(result, intList)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func CheckLevelSafety(levels []int) bool {
	// Safe if, the list is only increasing or decreasing
	// If 2 adjacent levels differ by at least 1 and at most 3
	prevLevel := levels[0]
	direction := 0
	for idx, level := range levels {
		if idx == 0 {
			continue
		}
		diff := level - prevLevel
		prevLevel = level
		if diff > 0 {
			if direction == -1 {
				return false
			}
			direction = 1
		} else if diff < 0 {
			if direction == 1 {
				return false
			}
			direction = -1
		}

		diff = int(math.Abs(float64(diff)))

		if diff < 1 || diff > 3 {
			return false
		}

	}

	return true
}
