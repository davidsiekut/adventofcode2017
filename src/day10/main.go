package main

import "fmt"

var lengths = []int{192, 69, 168, 160, 78, 1, 166, 28, 0, 83, 198, 2, 254, 255, 41, 12}

func main() {
	list := [256]int{}
	for i := 0; i < 256; i++ {
		list[i] = i
	}

	current := 0
	skip := 0
	for _, length := range lengths {
		remain := length
		curr := current
		s := []int{}
		for remain > 0 {
			s = append(s, list[curr])
			if curr >= len(list)-1 {
				curr = -1
			}
			curr++
			remain--
		}

		remain = length
		curr = current
		for remain > 0 {
			list[curr] = s[remain-1]
			if curr >= len(list)-1 {
				curr = -1
			}
			curr++
			remain--
		}

		current += length + skip
		if current > len(list) {
			current = current - len(list)
		}
		skip++
	}

	fmt.Println(list)
}
