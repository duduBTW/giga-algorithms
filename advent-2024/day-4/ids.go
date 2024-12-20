package day4

func generateId(indicator string) string {
	return "advent-of-code-day-4__" + indicator
}

func pt1(indicator string) string {
	return generateId("part-1__" + indicator)
}

var (
	IdData        = pt1("data")
	IdSearchForm  = pt1("serach-form")
	IdSearchInput = pt1("search-input")
	IdCanvas      = pt1("canvas")
	IdMatrix      = pt1("matrix")
	IdCell        = pt1("matrix-cell")
	IdFindX       = pt1("find-x")
	IdTotal       = pt1("total")
)

const Start = "AdventOfCodeDay4Part1Handler"
