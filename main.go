package main

import (
	"adventofcode/day08"
	"adventofcode/helpers"
	"fmt"
)

func main() {
	input := helpers.ReadInput("./day08/input.txt")
	fmt.Println(day08.Part1(input))
	fmt.Println(day08.Part2(input))
}
