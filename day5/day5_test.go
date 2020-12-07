package day5

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"testing"
)

func TestDayFiveInputOne(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	maxSeatID := 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		row := rowToInt(s[:7])
		col := colToInt(s[7:])
		seatID := row*8 + col
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	// output := DayFourPartOne(string(input), 1, 3)
	t.Log(fmt.Sprintf("%d\n", maxSeatID))
}

func TestDayFiveInputTwo(t *testing.T) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	seats := make([]int, 0)
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		row := rowToInt(s[:7])
		col := colToInt(s[7:])
		seatID := row*8 + col
		seats = append(seats, seatID)
	}
	sort.Ints(seats)
	var mySeat int
	var lastSeat int
	for _, s2 := range seats {
		if lastSeat == 0 {
			lastSeat = s2
			continue
		}
		if s2-lastSeat == 2 {
			mySeat = s2 - 1
			break
		} else {
			lastSeat = s2
		}
	}

	// output := DayFourPartOne(string(input), 1, 3)
	t.Log(fmt.Sprintf("%d\n", mySeat))
}
