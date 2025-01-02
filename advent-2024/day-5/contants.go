package day5

const Route = "/advent-2024/day-5"
const Start = "AdventOfCodeDay5Handler"

func id(indicator string) string {
	return "day-5__" + indicator
}

var (
	IdData       = id("data")
	IdTabs       = id("tabs")
	IdTabItem    = id("tab-item")
	IdFindButton = id("find-button")
	IdFixButton  = id("fix-button")
	IdTotal      = id("total")
	IdPages      = id("pages")
)

type PageStatus string

const (
	PageStatusUnknown PageStatus = "unknown"
	PageStatusValid   PageStatus = "valid"
	PageStatusInvalid PageStatus = "invalid"
)
