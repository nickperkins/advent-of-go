package day4

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
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
	return strings.Join(o, " "), true
}

func CheckRecord(input string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, v := range requiredFields {
		if !strings.Contains(input, v+":") {
			return false
		}
	}
	return true

}

func CheckRecord2(input string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, v := range requiredFields {
		if !strings.Contains(input, v+":") {
			return false
		}
	}
	fields := strings.Split(input, " ")
	for _, v := range fields {
		field, value := strings.Split(v, ":")[0], strings.Split(v, ":")[1]
		if !isValidField(field, value) {
			fmt.Printf("%v is %v", field, value)
			return false
		}

	}
	return true
}

func isValidField(field, value string) bool {
	switch field {
	case "byr":
		return byr(value)

	case "iyr":
		return iyr(value)

	case "eyr":
		return eyr(value)

	case "hgt":
		return hgt(value)

	case "hcl":
		return hcl(value)
	case "ecl":
		return ecl(value)
	case "pid":
		return pid(value)
	case "cid":
		return true

	}
	return false
}

func byr(value string) bool {
	v, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	if v >= 1920 && v <= 2002 {
		return true
	}
	return false
}

func iyr(value string) bool {
	v, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	if v >= 2010 && v <= 2020 {
		return true
	}
	return false
}

func eyr(value string) bool {
	v, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	if v >= 2020 && v <= 2030 {
		return true
	}
	return false
}

func hgt(value string) bool {
	unit := value[len(value)-2:]
	heightStr := value[0 : len(value)-2]
	height, err := strconv.Atoi(heightStr)
	if err != nil {
		return false
	}
	switch unit {
	case "in":
		if height >= 59 && height <= 76 {
			return true
		}
		break
	case "cm":
		if height >= 150 && height <= 193 {
			return true
		}
		break
	}
	return false
}

func ecl(value string) bool {
	validValues := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, v := range validValues {
		if v == value {
			return true
		}
	}
	return false
}

func hcl(value string) bool {
	match, err := regexp.Match(`(?i)#[0-9A-F]{6}`, []byte(value))
	if err != nil {
		return false
	}
	return match
}

func pid(value string) bool {
	if len(value) != 9 {
		return false
	}
	_, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return true
}
