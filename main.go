package main

import (
	"adventofcode/day11"
	"adventofcode/helpers"
	"fmt"
)

func main() {
	input := helpers.ReadInput("./day11/input.txt")
	fmt.Println(day11.Part1(input))
	fmt.Println(day11.Part2(input))
}
