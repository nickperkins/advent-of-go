package day1

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

	output := DayOnePartOne(string(input))
	t.Log(fmt.Sprintf("%v\n", output))
}

func TestDayTwoInput(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	output := DayOnePartTwo(string(input))
	t.Log(fmt.Sprintf("%v\n", output))
}
