package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
)

func main() {
    day1part2()
}
func day1part1() {
    calories := MostCaloriesPerElf("./input/day1.txt")

    fmt.Println(calories)
}


func day1part2() {
    calories := TopThreeCalories("./input/day1.txt")
    fmt.Println(calories)
}
func MostCaloriesPerElf(path string) int {
    lines := ReadInput(path)

    calories := GroupCaloriesByElf(lines)

    _, max := MinMax(calories)
    return max
}

func TopThreeCalories(path string) int {
    lines := ReadInput(path)
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

        if line == "" || index == len(lines) -1 {
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

func ReadInput(path string) []string {
    readFile, err := os.Open(path)

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    var lines []string
    for fileScanner.Scan() {
        lines = append(lines, fileScanner.Text())
    }

    return lines
}