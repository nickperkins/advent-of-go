package utils

import "strconv"

// MustAtoi will convert a string to an integer. It will panic if it can't be done.
func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
