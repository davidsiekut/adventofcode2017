package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	valid := 0
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		array := []string{}
		for _, a := range words {
			array = append(array, sortString(a))
		}
		if len(array) == len(trimDupes(array)) {
			valid++
		}
	}

	fmt.Println(valid)
}

type sorted []rune

func (s sorted) Len() int {
	return len(s)
}

func (s sorted) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sorted) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sorted(r))
	return string(r)
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
