package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var jumps []int
	for scanner.Scan() {
		s, _ := strconv.Atoi(scanner.Text())
		jumps = append(jumps, s)
	}

	curr := 0
	moves := 0
	for curr < len(jumps) {
		next := curr + jumps[curr]
		jumps[curr] = jumps[curr] + 1
		curr = next
		moves++
	}
	fmt.Println(moves)
}
