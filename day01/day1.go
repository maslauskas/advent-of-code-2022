package day01

import (
	"sort"
	"strconv"
)

func Part1(input []string) int {
	calories := GroupCaloriesByElf(input)
	sort.Ints(calories)
	last := calories[len(calories)-1:]
	return last[0]
}

func Part2(input []string) int {
	calories := GroupCaloriesByElf(input)

	last := calories[len(calories)-3:]
	sum := 0
	for _, cal := range last {
		sum += cal
	}

	return sum
}

func GroupCaloriesByElf(lines []string) []int {
	var calories []int
	cals := 0
	for index, line := range lines {
		value, _ := strconv.Atoi(line)
		cals += value

		if line == "" || index == len(lines)-1 {
			calories = append(calories, cals)
			cals = 0
		}
	}

	sort.Ints(calories)

	return calories
}
