package main

import (
	"adventofcode/day05"
	"adventofcode/helpers"
	"fmt"
)

func main() {
	input := helpers.ReadInput("./day05/input.txt")
	fmt.Println(day05.Part1(input))
	fmt.Println(day05.Part2(input))
}
