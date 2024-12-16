package day3

func generateId(indicator string) string {
	return "advent-of-code-day-3__" + indicator
}

func pt1(indicator string) string {
	return generateId("part-1__" + indicator)
}

var (
	IdsData               = pt1("data")
	CodeContent           = pt1("code-content")
	IdHilight             = pt1("hilight")
	IdsTotal              = pt1("total")
	IdsExpression         = pt1("expression")
	IdsExpressionMetadata = pt1("expression-metadata")
)
