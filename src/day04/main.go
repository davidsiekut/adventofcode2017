package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	valid := 0
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) == len(trimDupes(words)) {
			valid++
		}
	}

	fmt.Println(valid)
}

func trimDupes(s []string) []string {
	found := map[string]bool{}
	res := []string{}

	for x := range s {
		if found[s[x]] == true {
		} else {
			found[s[x]] = true
			res = append(res, s[x])
		}
	}
	return res
}
