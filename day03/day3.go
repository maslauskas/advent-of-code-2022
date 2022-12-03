package day03

import (
	"strings"
)

func Part1(input []string) int {
	score := 0
	for _, line := range input {
		item := FindCommonItem(line)
		score += GetScore(item)
	}

	return score
}

func Part2(input []string) int {
	score := 0
	groups := SplitGroups(input)
	for _, group := range groups {
		item := FindBadge(group)
		score += GetScore(item)
	}

	return score
}

func SplitItems(line string) []string {
	half := len(line) / 2

	first := string(line[0:(half)])
	second := string(line[half:])

	return []string{first, second}
}

func SplitGroups(input []string) [][]string {
	res := [][]string{}
	for i := 0; i < len(input); i = i + 3 {
		res = append(res, input[i:i+3])
	}

	return res
}

func FindCommonItem(input string) string {
	parts := SplitItems(input)

	first := parts[0]
	second := parts[1]

	for _, char := range first {
		index := strings.Index(second, string(char))
		if index != -1 {
			return string(second[index])
		}
	}
	return ""
}

func FindBadge(input []string) string {
	for _, char := range input[0] {
		index1 := strings.Index(input[1], string(char))
		index2 := strings.Index(input[2], string(char))

		if index1 != -1 && index2 != -1 {
			return string(char)
		}
	}
	return ""
}

func GetScore(item string) int {
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	score := strings.Index(alpha, item)
	return score + 1
}
