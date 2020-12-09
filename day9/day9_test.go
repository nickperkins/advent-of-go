package day9

import (
	"advent-of-go-2020/utils"
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

func TestDayNineInputOne(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := scanFile(t, input)
	dataInt := make([]int, len(data))
	for i, v := range data {
		dataInt[i] = utils.MustAtoi(v)
	}
	size := 25

	result := FindError(dataInt, size)

	fmt.Println(result)
}

func TestDayNineInputTwo(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := scanFile(t, input)
	dataInt := make([]int, len(data))
	for i, v := range data {
		dataInt[i] = utils.MustAtoi(v)
	}

	datapoint := 23278925
	result := FindWeakness(dataInt, datapoint)

	fmt.Println(result)
}
