package main

import (
	"bufio"
	"fmt"
	"os"
)

var input = ""
var commands = []string{}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = scanner.Text()
	}

	total := 0
	for i := 0; i < len(input); i++ {
		current := string(input[i])
		command := ""

		if current == "{" {
			depth := 1
			command += string(current)
			done := false
			score := 0
			for !done {
				i++
				current = string(input[i])
				if current == "!" {
					continue
				} else if current == "<" {
					end := false
					for !end {
						i++
						current = string(input[i])
						if current == "!" {
							i++
						} else if current == ">" {
							end = true
						}
					}
				} else if current == "}" {
					command += string(current)
					score += depth
					if depth == 1 {
						done = true
						break
					} else {
						depth--
					}
				} else if current == "{" {
					depth++
					command += string(current)
				}
			}
			total += score
		}
	}

	fmt.Println(total)
}
