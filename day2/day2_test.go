package day2

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestDayOneInput(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	output := DayTwoPartOne(string(input))
	t.Log(fmt.Sprintf("%v\n", output))
}

func TestDayOnePartTwoInput(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	output := DayTwoPartTwo(string(input))
	t.Log(fmt.Sprintf("%v\n", output))
}
