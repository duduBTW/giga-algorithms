package algorithms

type Algorithm struct {
	Name     string
	Link     string
	Category Category
}

type Category struct {
	Name        string
	Description string
	Icon        string
}

var Algorithms = []Algorithm{
	{
		Name:     "Day 1 - Historian Hysteria",
		Link:     "/advent-2024/day-1",
		Category: AdventOfCode2024,
	},
	{
		Name:     "Day 1 - Historian Hysteria Part 2",
		Link:     "/advent-2024/day-1-part-2",
		Category: AdventOfCode2024,
	},
	{
		Name:     "Day 2 - Red nose reports",
		Link:     "/advent-2024/day-2-part-1",
		Category: AdventOfCode2024,
	},
}

var AdventOfCode2024 = Category{
	Name:        "Advent of Code 2024",
	Description: "Advent of Code is an annual set of Christmas-themed computer programming challenges that follow an Advent calendar.",
	Icon:        "tree-pine",
}

var Categories = []Category{
	AdventOfCode2024,
	{
		Name:        "Sorting",
		Description: "Sorting is a fundamental operation in computer science that arranges elements of an array or list in a specific order.",
	},
	{
		Name:        "Searching",
		Description: "Searching is the process of finding a specific item or value within a collection of data.",
	},
	{
		Name:        "Dynamic Programming",
		Description: "Dynamic Programming is a technique used to solve problems by breaking them down into smaller subproblems and solving each subproblem only once.",
	},
}
