package day8

import (
	"fmt"
	"strconv"
	"strings"
)

var operations map[string]func(arg int)

type Operation struct {
	name     string
	argument int
	called   bool
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

func (c *CodeStack) RunCode() int {
	callStack := []int{}
	currLine := 0
	accum := 0
	for currLine < len(c.operations) {
		currOp := &c.operations[currLine]
		if !currOp.called {
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
			callStack = append(callStack, currLine)
			currLine += move
			currOp.called = true
		} else {
			fmt.Printf("Infinite loop on line %v - previous line is %v\n", currLine, callStack)
			return accum
		}

	}
	return accum
}
