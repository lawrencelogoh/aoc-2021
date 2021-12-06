package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var x int
var y int
var aim int

func main() {
	partOne()
	x, y = 0, 0
	partTwo()

}

func snc(s string) (string, int) {
	x := strings.Split(s, " ")
	a := x[0]
	b, err := strconv.Atoi(x[1])
	if err != nil {
		panic(err)
	}
	return a, b
}

func partOne() {
	fmt.Println("Part One")
	input, _ := ioutil.ReadFile("input.txt")
	list := strings.Split(string(input), "\n")
	for i := 0; i < len(list)-1; i++ {
		direction, degree := snc(list[i])
		if direction == "forward" {
			x = x + degree

		} else if direction == "down" {
			y = y + degree
		} else {
			y = y - degree
		}

	}
	fmt.Println(x * y)

}

func partTwo() {
	fmt.Println("Part Two")
	input, _ := ioutil.ReadFile("input.txt")
	list := strings.Split(string(input), "\n")
	for i := 0; i < len(list)-1; i++ {
		direction, degree := snc(list[i])
		if direction == "forward" {
			x = x + degree
			y = y + degree*aim

		} else if direction == "down" {
			aim = aim + degree
		} else {
			aim = aim - degree
		}

	}
	fmt.Println(x * y)

}
