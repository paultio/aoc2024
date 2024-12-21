package aoc

import (
	"fmt"
	"strings"
)

type Rule struct {
	Before int
	After  int
}

func ParseRule(line string) Rule {
	var before, after int
	fmt.Sscanf(line, "%d|%d", &before, &after)
	return Rule{before, after}
}

func (r Rule) Apply(pages []int) bool {
	beforeIdx := -1
	afterIdx := -1
	for i, page := range pages {
		if page == r.Before {
			beforeIdx = i
		}
		if page == r.After {
			afterIdx = i
		}
	}

	if beforeIdx == -1 || afterIdx == -1 {
		return true
	}

	return beforeIdx < afterIdx
}

func ParseFromStr(content string) (rules []Rule, orderings [][]int) {
	cnt := strings.Split(content, "\n\n")

	for _, line := range strings.Split(cnt[0], "\n") {
		rules = append(rules, ParseRule(line))
	}

	for _, line := range strings.Split(cnt[1], "\n") {
		ord := strings.Split(line, ",")
		ordering := make([]int, len(ord))
		for i, o := range ord {
			fmt.Sscanf(o, "%d", &ordering[i])
		}
		orderings = append(orderings, ordering)
	}
	return rules, orderings
}

func ApplyRules(rules []Rule, pages [][]int) (correct_idx []int) {
	for idx, page := range pages {
		if len(page)%2 == 0 {
			fmt.Println("Page has even number of pages: ", page)
		}
		ok := true
		for _, rule := range rules {
			if !rule.Apply(page) {
				fmt.Println("Rule failed: ", rule, " for page: ", page)
				ok = false
				break
			}
		}
		if ok {
			correct_idx = append(correct_idx, idx)
		}
	}
	return correct_idx
}
