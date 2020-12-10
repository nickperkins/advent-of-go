package day10

import (
	"sort"
)

func sortAdapters(i []int) []int {
	// add 0 to account for the wall socket
	i = append(i, 0)
	// sort in order
	sort.Ints(i)
	// append one +3 of the highest value
	i = append(i, i[len(i)-1]+3)

	return i
}
func GetDifferences(input []int) (int, int) {

	var ones, threes int

	sorted := sortAdapters(input)

	for i := 0; i < len(sorted)-1; i++ {
		difference := sorted[i+1] - sorted[i]
		switch difference {
		case 1:
			ones++
		case 3:
			threes++
		}
	}

	return ones, threes
}

func isValid(input []int) bool {

	for i := 0; i < len(input)-1; i++ {
		difference := input[i+1] - input[i]
		if difference > 0 && difference <= 3 {
			continue
		} else {
			return false
		}

	}
	return true

}
