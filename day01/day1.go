package day01

import (
	"adventofcode/helpers"
	"fmt"
	"sort"
	"strconv"
)

func Part1(path string) {
	calories := MostCaloriesPerElf(path)

	fmt.Println(calories)
}

func Part2(path string) {
	calories := TopThreeCalories(path)
	fmt.Println(calories)
}
func MostCaloriesPerElf(path string) int {
	lines := helpers.ReadInput(path)

	calories := GroupCaloriesByElf(lines)

	_, max := MinMax(calories)
	return max
}

func TopThreeCalories(path string) int {
	lines := helpers.ReadInput(path)
	calories := GroupCaloriesByElf(lines)

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

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
