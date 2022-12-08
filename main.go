package main

import (
	"adventofcode/day07"
	"adventofcode/helpers"
	"fmt"
)

func main() {
	input := helpers.ReadInput("./day07/input.txt")
	fmt.Println(day07.Part1(input))
	fmt.Println(day07.Part2(input))
}
