package main

import (
	"adventofcode/day01"
	"adventofcode/helpers"
	"fmt"
)

func main() {
	input := helpers.ReadInput("./day01/input.txt")
	fmt.Println(day01.Part1(input))
	fmt.Println(day01.Part2(input))
}
