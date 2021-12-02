package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var nums []int = input()

func main() {
	partOne()
	partTwo()

}

func partOne() {
	fmt.Println("Part one")
	x := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			x++
		}

	}
	fmt.Println(x)

}
func partTwo() {
	fmt.Println("Part two")
	x := 0
	for i := 0; i < len(nums)-3; i++ {
		a := nums[i] + nums[i+1] + nums[i+2]
		b := nums[i+1] + nums[i+2] + nums[i+3]
		if b > a {
			x++
		}

	}
	fmt.Println(x)

}

func input() []int {
	input, _ := ioutil.ReadFile("input.txt")
	n := []int{}

	list := strings.Split(string(input), "\n")

	for i := 0; i < len(list)-1; i++ {
		j, err := strconv.Atoi(list[i])
		if err != nil {
			panic(err)
		}
		n = append(n, j)
	}
	return n
}
