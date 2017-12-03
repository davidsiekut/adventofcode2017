package main

import "fmt"
import "math"

const (
	target = 361527
)

func main() {
	current := 1
	moves := 1
	remaining := 1
	c := 2
	direction := 0

	x := 0
	y := 0

	for current < target {
		for remaining > 0 {
			switch direction {
			case 0:
				x++
			case 1:
				y--
			case 2:
				x--
			case 3:
				y++
			}
			current++
			if current == target {
				fmt.Printf("%f\n", math.Abs(float64(x))+math.Abs(float64(y)))
			}
			remaining--
		}

		direction++
		if direction > 3 {
			direction = 0
		}

		c--
		if c == 0 {
			c = 2
			moves++
		}
		remaining = moves
	}
}
