package day12

type PathHeap []Path

type Path struct {
	Score  int
	Coords []string
}

func (h PathHeap) Len() int {
	return len(h)
}
func (h PathHeap) Less(i, j int) bool {
	return h[i].Score < h[j].Score
}
func (h PathHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PathHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Path))
}

func (h *PathHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

type HeightMap struct {
	Rows  []string
	Start [2]int
	End   [2]int
}

func CreateHeightMap(input []string) HeightMap {
	// find start
	var start [2]int
	var end [2]int
	for y, row := range input {
		for x, char := range row {
			if string(char) == "S" {
				start = [2]int{y, x}
			}

			if string(char) == "E" {
				end = [2]int{y, x}
			}
		}
	}

	return HeightMap{
		Rows:  input,
		Start: start,
		End:   end,
	}
}
