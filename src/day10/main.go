package main

import "fmt"

func main() {
	list := [256]int{}
	for i := 0; i < 256; i++ {
		list[i] = i
	}

	rounds := 64
	current := 0
	skip := 0
	input := "192,69,168,160,78,1,166,28,0,83,198,2,254,255,41,12"
	lengths := []int{}
	for _, r := range input {
		lengths = append(lengths, int(r))
	}
	app := []int{17, 31, 73, 47, 23}
	for _, a := range app {
		lengths = append(lengths, a)
	}

	for rounds > 0 {
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

			left := length + skip
			for left > 0 {
				current++
				if current >= len(list) {
					current = 0
				}
				left--
			}

			skip++
		}
		rounds--
	}

	hash := []int{}
	for i := 0; i < 16; i++ {
		xor := list[i*16]
		for j := 1; j < 16; j++ {
			xor = xor ^ list[(i*16)+j]
		}
		hash = append(hash, xor)
	}

	s := ""
	for _, h := range hash {
		s += fmt.Sprintf("%02x", h)
	}
	fmt.Println(s)
}
