package advent2024day2

func generateId(indicator string) string {
	return "advent-of-code-day__" + indicator
}

func pt1(indicator string) string {
	return generateId("part-1__" + indicator)
}

var (
	IdsPart1TotalContainer = pt1("total-container")
	IdsPart1TotalValue     = pt1("total-value")
	IdsCalculateAll        = pt1("calculate-total-button")
	IdsData                = pt1("data")
	IdsPt1List             = pt1("list")
	IdsPt1RadioContainer   = pt1("radio-container")
	IdsPt1Radio            = pt1("radio")
)
