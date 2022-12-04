package main

import (
	"adventofcode/day04"
	"adventofcode/helpers"
	"fmt"
)

func main() {
	input := helpers.ReadInput("./day04/input.txt")
	fmt.Println(day04.Part1(input))
	fmt.Println(day04.Part2(input))
}
