package day3

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestDayThreeInput(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	output := DayThreePartOne(string(input), 1, 3)
	t.Log(fmt.Sprintf("%v\n", output))
}

func TestDayThreePartTwoInput(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	output := DayThreePartOne(string(input), 1, 1)
	output *= DayThreePartOne(string(input), 1, 3)
	output *= DayThreePartOne(string(input), 1, 5)
	output *= DayThreePartOne(string(input), 1, 7)
	output *= DayThreePartOne(string(input), 2, 1)
	t.Log(fmt.Sprintf("%v\n", output))
}
