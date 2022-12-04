package day04

import (
	"strconv"
	"strings"
)

func Part1(input []string) int {
	score := 0

	for _, row := range input {
		s1, s2 := MakeSections(row)
		if s1.Overlaps(s2) {
			score++
		}
	}

	return score
}

func Part2(input []string) int {
	score := 0

	for _, row := range input {
		s1, s2 := MakeSections(row)
		if s1.Intersects(s2) {
			score++
		}
	}

	return score
}

type Section struct {
	Start int
	End   int
}

func (s1 Section) Overlaps(s2 Section) bool {
	return s1.Contains(s2) || s2.Contains(s1)
}

func (s1 Section) Contains(s2 Section) bool {
	// will not contain if s1 starts after s2
	if s1.Start > s2.Start {
		return false
	}

	// will not contain if s1 ends before s2
	if s1.End < s2.End {
		return false
	}

	return true
}

func (s1 Section) Intersects(s2 Section) bool {
	// will not intersect if s1 starts after s2 ends
	if s1.Start > s2.End {
		return false
	}

	// will not intersect if s2 starts after s1 ends
	if s2.Start > s1.End {
		return false
	}
	return true
}

func MakeSection(input string) Section {
	parts := strings.Split(input, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	return Section{start, end}
}

func MakeSections(input string) (Section, Section) {
	parts := strings.Split(input, ",")
	s1 := MakeSection(parts[0])
	s2 := MakeSection(parts[1])

	return s1, s2
}
