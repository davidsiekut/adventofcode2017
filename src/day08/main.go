package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type register struct {
	ID    string
	Value int
}

var commands = []string{}
var registers = []register{}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	max := 0
	for _, cmd := range commands {
		split := strings.Split(cmd, " ")

		reg := split[0]
		per := split[1]
		amt, _ := strconv.Atoi(split[2])
		cmp := split[4]
		cnd := split[5]
		val, _ := strconv.Atoi(split[6])

		if !contains(registers, reg) {
			registers = append(registers, register{ID: reg, Value: 0})
		}

		if !contains(registers, cmp) {
			registers = append(registers, register{ID: cmp, Value: 0})
		}

		mult := 1
		if per == "dec" {
			mult = -1
		}

		v := &registers[pos(registers, reg)]
		w := registers[pos(registers, cmp)]
		if cnd == ">" {
			if w.Value > val {
				v.Value += mult * amt
			}
		} else if cnd == "<" {
			if w.Value < val {
				v.Value += mult * amt
			}
		} else if cnd == "<=" {
			if w.Value <= val {
				v.Value += mult * amt
			}
		} else if cnd == ">=" {
			if w.Value >= val {
				v.Value += mult * amt
			}
		} else if cnd == "==" {
			if w.Value == val {
				v.Value += mult * amt
			}
		} else if cnd == "!=" {
			if w.Value != val {
				v.Value += mult * amt
			}
		} else {
			fmt.Println(cnd)
		}

		if v.Value > max {
			max = v.Value
		}
	}

	fmt.Println(max)
}

func contains(array []register, key string) bool {
	for _, reg := range array {
		if reg.ID == key {
			return true
		}
	}
	return false
}

func pos(array []register, key string) int {
	for pos, reg := range array {
		if reg.ID == key {
			return pos
		}
	}
	return -1
}
