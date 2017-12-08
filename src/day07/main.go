package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Disc is a disc
type Disc struct {
	ID       string
	Weight   int
	Children []Disc
}

var discs = []Disc{}
var mem = []string{}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mem = append(mem, scanner.Text())
	}

	//disc := make("hlqnsbe")
	//disc := make("rilyl")
	//disc := make("aurik")
	disc := make("jriph")
	for _, child := range disc.Children {
		fmt.Print(child.ID, " ", child.Weight, " ")
	}
	// 2102 - 1998 = 5
	// 1998 - 5 = 1993
}

func make(id string) Disc {
	for _, a := range mem {
		split := strings.Split(a, " ")

		if split[0] == id {
			s := strings.Replace(split[1], "(", "", -1)
			s = strings.Replace(s, ")", "", -1)
			weight, _ := strconv.Atoi(s)

			children := []Disc{}
			for i := 3; i < len(split); i++ {
				c := strings.Split(split[i], ",")[0]
				child := make(c)
				children = append(children, child)
				weight += child.Weight
			}

			disc := Disc{
				ID:       split[0],
				Weight:   weight,
				Children: children,
			}

			return disc
		}
	}
	return Disc{}
}
