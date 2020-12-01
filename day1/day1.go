package day1

import (
	"sort"
	"strconv"
	"strings"
)

func convertToSortedSlice(input string) []int {
	entries := []int{}

	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		i, _ := strconv.Atoi(s)
		entries = append(entries, i)
	}
	// Sort the slice
	sort.Ints(entries)
	return entries
}

// DayOnePartOne is the solution for part one
func DayOnePartOne(input string) int {
	match := 2020

	entries := convertToSortedSlice(input)

	for i, temp1 := range entries {
		for _, temp2 := range entries[i+1:] {
			if temp1+temp2 == match {
				return temp1 * temp2
			}
		}
	}
	return 0
}

// DayOnePartTwo is the solution for part two
func DayOnePartTwo(input string) int {
	match := 2020

	entries := convertToSortedSlice(input)

	for i, temp1 := range entries {
		for j, temp2 := range entries[i+1:] {
			for _, temp3 := range entries[j+1:] {
				if temp1+temp2+temp3 == match {
					return temp1 * temp2 * temp3
				}
			}
		}

	}
	return 0
}
