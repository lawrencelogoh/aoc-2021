package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var eRate string
var gRate string

// var power int = gRate * eRate

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("Part One")
	input, _ := ioutil.ReadFile("input.txt")
	list := strings.Split(string(input), "\n")
	l := strings.Split(string(list[0]), "") // 00100

	for j := 0; j < len(l); j++ {
		n := []string{}
		for i := 0; i < len(list)-1; i++ {
			l := strings.Split(string(list[i]), "") // 00100

			n = append(n, l[j])

		}
		numbers := strings.Join(n, "")
		cb(numbers)

	}
	g, _ := strconv.ParseInt(gRate, 2, 64)
	e, _ := strconv.ParseInt(eRate, 2, 64)
	fmt.Println(g * e)

}

func partTwo() {
	fmt.Println("Part Two")

}

func cb(s string) {
	z := strings.Count(s, "0")
	o := strings.Count(s, "1")

	if z > o {
		gRate = gRate + "0"
		eRate = eRate + "1"
	} else {
		gRate = gRate + "1"
		eRate = eRate + "0"
	}
}
