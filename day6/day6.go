package day6

import (
	"bufio"
	"strings"
)

func ReadBlock(scanner *bufio.Scanner) (string, bool) {
	var o []string
	if scanner.Scan() == false {
		return "", false
	}

	for len(scanner.Text()) > 0 {
		o = append(o, scanner.Text())
		if scanner.Scan() == false {
			break
		}
	}
	return strings.Join(o, ""), true
}

func ReadBlockSplit(scanner *bufio.Scanner) (string, bool) {
	var o []string
	if scanner.Scan() == false {
		return "", false
	}

	for len(scanner.Text()) > 0 {
		o = append(o, scanner.Text())
		if scanner.Scan() == false {
			break
		}
	}
	return strings.Join(o, ","), true
}

func uniqueLetters(input string) int {
	letterMap := make(map[rune]struct{})
	// trim spaces
	input = strings.TrimSpace(input)
	for _, s := range input {
		letterMap[s] = struct{}{}
	}

	return len(letterMap)
}

func sameAnswers(input string) int {
	letterMap := make(map[rune]int)
	// trim spaces
	input = strings.TrimSpace(input)

	people := strings.Split(input, ",")
	var peopleCount int
	var validAnswers []string
	for _, p := range people {
		peopleCount++
		for _, s := range p {

			letterMap[s]++

		}

	}
	for k, v := range letterMap {
		if v == peopleCount {
			validAnswers = append(validAnswers, string(k))
		}
	}

	return len(validAnswers)
}
