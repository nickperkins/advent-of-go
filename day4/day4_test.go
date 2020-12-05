package day4

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestDayThreeInput(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	count := 0
	for {
		l, ok := ReadBlock(scanner)
		if !ok {
			break
		}
		if CheckRecord(l) {
			count++
		}
	}

	// output := DayFourPartOne(string(input), 1, 3)
	t.Log(fmt.Sprintf("%d\n", count))
}

func TestDayFourInputTwo(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	count := 0
	for {
		l, ok := ReadBlock(scanner)
		if !ok {
			break
		}
		if CheckRecord2(l) {
			count++
		}
	}

	// output := DayFourPartOne(string(input), 1, 3)
	t.Log(fmt.Sprintf("%d\n", count))
}
