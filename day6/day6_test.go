package day6

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestDaySixInputOne(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	sum := 0
	for {
		l, ok := ReadBlock(scanner)
		if !ok {
			break
		}
		sum += uniqueLetters(l)
	}

	// output := DayFourPartOne(string(input), 1, 3)
	t.Log(fmt.Sprintf("%d\n", sum))
}

func TestDaySixInputTwo(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	sum := 0
	for {
		l, ok := ReadBlockSplit(scanner)
		if !ok {
			break
		}

		sum += sameAnswers(l)
	}

	// output := DayFourPartOne(string(input), 1, 3)
	t.Log(fmt.Sprintf("%d\n", sum))
}
