package main

import "fmt"

const (
	target = 361527
)

// Point is a point
type Point struct {
	x, y, val int
}

var mem = []Point{Point{x: 0, y: 0, val: 1}}

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

			sum := 0
			sum += find(x-1, y-1)
			sum += find(x-1, y+1)
			sum += find(x+1, y-1)
			sum += find(x+1, y+1)
			sum += find(x, y-1)
			sum += find(x, y+1)
			sum += find(x-1, y)
			sum += find(x+1, y)
			current = sum

			mem = append(mem, Point{x: x, y: y, val: current})

			if current > target {
				fmt.Printf("%d\n", current)
				break
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

func find(x int, y int) int {
	for _, v := range mem {
		if v.x == x {
			if v.y == y {
				return v.val
			}
		}
	}
	return 0
}
