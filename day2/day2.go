package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const regexStr = `(\d+)-(\d+) (\w{1}): (\w+)`

func convertToSlice(input string) []string {
	entries := []string{}

	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		entries = append(entries, s)
	}

	return entries
}

func DayTwoPartOne(input string) int {
	s := convertToSlice(input)
	re := regexp.MustCompile(regexStr)
	valid := 0
	for _, v := range s {
		matchGroups := re.FindStringSubmatch(v)
		low, _ := strconv.Atoi(matchGroups[1])
		high, _ := strconv.Atoi(matchGroups[2])
		charCount := strings.Count(matchGroups[4], matchGroups[3])
		fmt.Printf("%v %v %v\n", low, high, charCount)
		if charCount >= low && charCount <= high {
			valid++
			fmt.Println("Valid")
		}
	}

	return valid
}

func DayTwoPartTwo(input string) int {
	s := convertToSlice(input)
	re := regexp.MustCompile(regexStr)
	valid := 0
	for _, v := range s {
		matchGroups := re.FindStringSubmatch(v)
		pos1, _ := strconv.Atoi(matchGroups[1])
		pos2, _ := strconv.Atoi(matchGroups[2])
		match1 := string(matchGroups[4][pos1-1])
		match2 := string(matchGroups[4][pos2-1])
		//charCount := strings.Count(matchGroups[4], matchGroups[3])
		//fmt.Printf("%v %v %v\n", low, high, charCount)
		if match1 != match2 && (match1 == matchGroups[3] || match2 == matchGroups[3]) {
			valid++
			fmt.Println("Valid")
		}
	}

	return valid
}
