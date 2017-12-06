package main

import (
	"fmt"
	"strconv"
)

func main() {
	bank := [16]int{2, 8, 8, 5, 4, 2, 3, 1, 5, 5, 1, 2, 15, 13, 5, 14}
	mem := []string{"2885423155121513514"}
	found := 0
	moves := 0

	for found != 2 {
		max := 0
		pos := 0
		for p, e := range bank {
			if e > max {
				max = e
				pos = p
			}
		}

		distribute := bank[pos]
		bank[pos] = 0
		pos++
		for distribute > 0 {
			if pos == 16 {
				pos = 0
			}
			bank[pos] = bank[pos] + 1
			distribute--
			pos++
		}

		str := ""
		for _, a := range bank {
			str += strconv.Itoa(a)
		}

		found = 0
		for _, a := range mem {
			if a == str {
				found++
			}
		}

		mem = append(mem, str)
		moves++
	}

	fmt.Println(moves)
}
