package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type Bag struct {
	style  string
	colour string
}

type BagRules map[Bag]map[Bag]int

var parsedRules BagRules

func parseRules(input []string) BagRules {
	parsedRules := BagRules{}

	for _, v := range input {
		splitValue := strings.Split(v, " contain ")
		splitBagData := strings.Split(splitValue[0], " ")
		thisBag := Bag{
			style:  splitBagData[0],
			colour: splitBagData[1],
		}

		parsedRules[thisBag] = map[Bag]int{}

		splitContents := strings.Split(splitValue[1], ", ")
		for _, v := range splitContents {
			if !strings.Contains(v, "other bags.") {
				splitValue := strings.Split(v, " ")
				parsedRules[thisBag][Bag{
					style:  splitValue[1],
					colour: splitValue[2],
				}], _ = strconv.Atoi(splitValue[0])

			}
		}
	}
	return parsedRules

}

func (b *Bag) canContain(rules BagRules, bag Bag) bool {
	bagContains := rules[*b]
	for k := range bagContains {
		if k == bag {
			return true
		}
		if k.canContain(rules, bag) {
			return true
		}

	}
	return false
}

func (b *Bag) getCost(rules BagRules, depth int) int {
	fmt.Println(depth)
	count := 0
	bagContains := rules[*b]
	// how many bags go with this bag?
	// get how many bags this bag contains
	// multiply by the depth
	for k, v := range bagContains {

		count += k.getCost(rules, depth+1) * v
		count += v

	}
	return count
}
