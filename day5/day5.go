package day5

import (
	"strconv"
	"strings"
)

func rowToInt(input string) int {
	input = strings.ReplaceAll(input, "F", "0")
	input = strings.ReplaceAll(input, "B", "1")

	result, err := strconv.ParseInt(input, 2, 0)

	if err != nil {
		panic(err)
	}

	return int(result)
}

func colToInt(input string) int {
	input = strings.ReplaceAll(input, "L", "0")
	input = strings.ReplaceAll(input, "R", "1")

	result, err := strconv.ParseInt(input, 2, 0)

	if err != nil {
		panic(err)
	}

	return int(result)
}
