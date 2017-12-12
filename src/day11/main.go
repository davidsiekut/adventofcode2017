package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var input = ""

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = scanner.Text()
	}

	path := strings.Split(input, ",")

	x := 0
	y := 0
	z := 0
	max := 0
	for _, mov := range path {
		if mov == "n" {
			x++
			z--
		} else if mov == "ne" {
			y++
			z--
		} else if mov == "se" {
			x--
			y++
		} else if mov == "s" {
			x--
			z++
		} else if mov == "sw" {
			y--
			z++
		} else if mov == "nw" {
			x++
			y--
		}
		m := int(distance(0, 0, 0, x, y, z))
		if m > max {
			max = m
		}
	}

	fmt.Println(max)
}

func distance(x1, y1, z1, x2, y2, z2 int) float64 {
	max := math.Max(math.Abs(float64(x1-x2)), math.Abs(float64(y1-y2)))
	return math.Max(max, math.Abs(float64(z1-z2)))
}
