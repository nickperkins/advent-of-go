package day7

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func scanFile(t *testing.T, input []byte) []string {
	t.Helper()
	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	var o []string
	if scanner.Scan() == false {
		return o
	}

	for len(scanner.Text()) > 0 {
		o = append(o, scanner.Text())
		if scanner.Scan() == false {
			break
		}
	}
	return o

}

func TestDaySevenInputOne(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := scanFile(t, input)
	rules := parseRules(data)

	testBag := Bag{
		style:  "shiny",
		colour: "gold",
	}

	count := 0
	for k := range rules {
		if k.canContain(rules, testBag) {
			count++
		}
	}
	fmt.Println(count)
}

func TestDaySevenInputTwo(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := scanFile(t, input)
	rules := parseRules(data)

	testBag := Bag{
		style:  "shiny",
		colour: "gold",
	}

	cost := testBag.getCost(rules, 1)

	fmt.Println(cost)
}
