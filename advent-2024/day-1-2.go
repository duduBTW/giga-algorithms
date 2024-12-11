package advent2024

func NewDay12LookupTable(values []int) map[int]int {
	lookup := make(map[int]int)
	for _, value := range values {
		lookup[value]++
	}
	return lookup
}

func Day12CalculateTotal(lookup map[int]int, list []int) int {
	total := 0
	for index, _ := range list {
		total += Day12Line(lookup, list, index)
	}
	return total
}

func Day12Line(lookup map[int]int, list []int, index int) int {
	currentItem := list[index]
	currentApparences := lookup[currentItem]
	return currentItem * currentApparences
}
