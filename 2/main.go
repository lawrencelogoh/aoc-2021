package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var x int
var y int

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	list := strings.Split(string(input), "\n")
	for i := 0; i < len(list)-1; i++ {
		direction, depth := parse(list[i])
		if direction == "forward" {
			x = x + depth

		} else if direction == "down" {
			y = y + depth
		} else {
			y = y - depth
		}

	}
	fmt.Println(x * y)

}

func parse(s string) (string, int) {
	x := strings.Split(s, " ")
	a := x[0]
	b, err := strconv.Atoi(x[1])
	if err != nil {
		panic(err)
	}
	return a, b
}
