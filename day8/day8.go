package day8

import (
	"errors"
	"strconv"
	"strings"
)

var operations map[string]func(arg int)

type Operation struct {
	name     string
	argument int
}

type CodeStack struct {
	operations []Operation
}

func (c *CodeStack) LoadCode(source []string) {
	for _, v := range source {
		op, argStr := strings.Split(v, " ")[0], strings.Split(v, " ")[1]
		arg, _ := strconv.Atoi(argStr)
		c.operations = append(c.operations,
			Operation{
				name:     op,
				argument: arg,
			})
	}

}

func (c *CodeStack) RunCode() (int, error) {
	callStack := []int{}
	currLine := 0
	accum := 0
	for currLine < len(c.operations) {
		currOp := &c.operations[currLine]
		for _, i := range callStack {
			if i == currLine {
				return accum, errors.New("Infinite Loop")
			}
		}
		callStack = append(callStack, currLine)

		move := 0
		switch currOp.name {
		case "nop":
			move++
			break
		case "jmp":
			move = currOp.argument
			break
		case "acc":
			accum += currOp.argument
			move++
			break
		}

		currLine += move

	}
	return accum, nil
}

func (c *CodeStack) FixCode() int {
	for i := range c.operations {
		currop := c.operations[i].name
		switch currop {
		case "jmp":
			c.operations[i].name = "nop"
		case "nop":
			c.operations[i].name = "jmp"
		case "acc":
			continue
		}

		result, err := c.RunCode()
		if err == nil {
			return result

		}
		c.operations[i].name = currop

	}
	panic("No solution")

}
