package main

import (
	"adventofcode/day02"
	"adventofcode/helpers"
	"fmt"
)

func main() {
	input := helpers.ReadInput("./day02/input.txt")
	fmt.Println(day02.Part1(input))
	fmt.Println(day02.Part2(input))
}
