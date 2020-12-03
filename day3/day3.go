package day3

import (
	"strings"
)

const regexStr = `(\d+)-(\d+) (\w{1}): (\w+)`

func convertToGrid(input string) [][]string {

	splitStrings := strings.Split(strings.TrimSpace(string(input)), "\n")
	entries := make([][]string, len(splitStrings))
	for l, s := range splitStrings {

		items := strings.Split(s, "")
		entries[l] = make([]string, len(items))
		for k, v := range items {
			entries[l][k] = v
		}
	}

	return entries
}

func DayThreePartOne(input string, down, right int) int {
	s := convertToGrid(input)
	var trees int
	// get the row and column count
	rows := len(s)
	cols := len(s[0])

	j := 0
	for i := down; i <= rows-1; i += down {
		j += right
		if j >= cols {

			j -= cols

		}
		if s[i][j] == "#" {
			//	fmt.Println("Ouch!")
			trees++
		} else {
			//	fmt.Println("Phew!")
		}
	}

	return trees
}
