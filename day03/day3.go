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

func SplitItems(line string) []string {
	half := len(line) / 2

	first := string(line[0:(half)])
	second := string(line[half:])

	return []string{first, second}
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

func GetScore(item string) int {
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	score := strings.Index(alpha, item)
	return score + 1
}
