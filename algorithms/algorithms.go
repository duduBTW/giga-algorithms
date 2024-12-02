package algorithms

type Algorithm struct {
	Name string
	Category    Category
}

type Category struct {
	Name string
	Description string
}

var Algorithms = []Algorithm{
	{
		Name:        "Day 1 - Historian Hysteria",
		Category:    AdventOfCode2024,
	},
}

var AdventOfCode2024 = Category{
	Name: "Advent of Code 2024",
	Description: "Advent of Code is an annual set of Christmas-themed computer programming challenges that follow an Advent calendar.",
}

var Categories = []Category{
	AdventOfCode2024,
	{
		Name: "Sorting",
		Description: "Sorting is a fundamental operation in computer science that arranges elements of an array or list in a specific order.",
	},
	{
		Name: "Searching",
		Description: "Searching is the process of finding a specific item or value within a collection of data.",
	},
	{
		Name: "Dynamic Programming",
		Description: "Dynamic Programming is a technique used to solve problems by breaking them down into smaller subproblems and solving each subproblem only once.",
	},
}
