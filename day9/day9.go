package day9

import "sort"

func FindError(input []int, size int) int {
	var pointer int
	var testvalue int

label:
	for pointer = size; pointer < len(input); pointer++ {
		preamble := input[pointer-size : pointer]
		testvalue = input[pointer]
		for i := 0; i < size; i++ {
			for j := i + 1; j < size; j++ {
				if preamble[i]+preamble[j] == testvalue {
					continue label
				}

			}

		}
		return testvalue
	}

	panic("Error not found")
}

func FindWeakness(input []int, target int) int {
	var pointer int

	for pointer = 0; pointer < len(input); pointer++ {
		testsum := 0
		for i := 0; ; i++ {
			testsum += input[pointer+i]
			if testsum > target {
				break
			}
			if testsum == target {
				inputs := input[pointer : pointer+i]
				sort.Ints(inputs)
				return inputs[0] + inputs[len(inputs)-1]
			}
		}

	}
	return 0
}
