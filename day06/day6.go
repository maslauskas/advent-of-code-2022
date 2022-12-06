package day06

func Part1(input []string) int {
	return FindMarker(input[0])
}

func FindMarker(packet string) int {
	for i := 0; i < len(packet)-4; i++ {
		marker := packet[i : i+4]
		duplicates := make(map[int32]int32, 4)

		for _, mark := range marker {
			if _, exists := duplicates[mark]; !exists {
				duplicates[mark] = mark
			}
		}

		if len(duplicates) == len(marker) {
			return i + 4
		}
	}

	return -1
}
