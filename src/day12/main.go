package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type program struct {
	Links []int
}

var input = ""
var programs = [2000]program{}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = scanner.Text()
		split := strings.Split(input, " ")

		id, _ := strconv.Atoi(split[0])
		links := []int{}
		for i := 2; i < len(split); i++ {
			a := strings.Split(split[i], ",")
			b, _ := strconv.Atoi(a[0])
			links = append(links, b)
		}

		programs[id] = program{Links: links}
	}

	x := 0
	for pos := range programs {
		b := hasLink(pos, 0, []int{pos})
		if b {
			x++
		}
	}

	fmt.Println(x)
}

func hasLink(id int, link int, visited []int) bool {
	for _, prog := range programs[id].Links {
		if prog == link {
			return true
		}
		if !contains(visited, prog) {
			visited = append(visited, prog)
			b := hasLink(prog, link, visited)
			if b {
				return true
			}
		}
	}
	return false
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
