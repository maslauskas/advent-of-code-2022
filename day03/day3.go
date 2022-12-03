package day03

func SplitItems(line string) []string {
	half := len(line) / 2

	first := string(line[0:(half)])
	second := string(line[half:])

	return []string{first, second}
}
