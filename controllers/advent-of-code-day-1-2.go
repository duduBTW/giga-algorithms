//go:build js && wasm
// +build js,wasm

package controllers

import (
	"fmt"
	"syscall/js"

	"github.com/dudubtw/giga-algorithms/components"
	"github.com/dudubtw/giga-algorithms/constants"
	"github.com/dudubtw/giga-algorithms/js-layer"
	"github.com/dudubtw/giga-algorithms/advent-2024"
	"github.com/a-h/templ"
)

func getJsonData2() [][]int {
	fmt.Println("id", constants.Advent2024Day1Part2LookupTableDataID)
	newData, err := jslayer.GetJsonData(constants.Advent2024Day1Part2LookupTableDataID)
	if err != nil {
		fmt.Println("Error getting json data", err)
		return nil
	}
	return newData
}

var currentDay1pt2DataIndex = 0

var totalDay1pt2 = StateProps[int]{
	Value: 0,
	Target: jslayer.Id(constants.Advent2024Day1Part2TotalID),
	RenderComponent: func(value int) templ.Component {
		return components.Day1Pt2TotalValue(value)
	},
}

var day1pt2Data = StateProps[[][]int]{
	Value: make([][]int, 2),
	Target: jslayer.Id(constants.Advent2024Day1Part2ListID),
	RenderComponent: func(value [][]int) templ.Component {
		return components.Day2List(value[0])
	},
}
var generateLookupTableHandler = jslayer.EventListener{
	Selector:    jslayer.Id(constants.Advent2024Day1Part2GenerateLookupTableID),
	EventType:   "click",
	Listener: func(this js.Value, args []js.Value) {
		lookupTable.Set(advent2024.NewDay12LookupTable(day1pt2Data.Value[1]))
	},
}

var lookupTable = StateProps[map[int]int]{
	Value: make(map[int]int),
	Target: jslayer.Id(constants.Advent2024Day1Part2LookupTableID),
	RenderComponent: func(value map[int]int) templ.Component {
		return components.Day2LookupTable(value)
	},
}

var nextDay1pt2ClickHandler = jslayer.EventListener{
	Selector:  jslayer.Id(constants.Advent2024Day1Part2NextID),
	EventType: "click",
	Listener: func(this js.Value, args []js.Value) {
		fmt.Println("nextClickHandler", lookupTable.Value)
		if len(lookupTable.Value) == 0 {
			fmt.Println("No lookup table")
			return
		}

		totalDay1pt2.Set(advent2024.Day12Line(lookupTable.Value, day1pt2Data.Value[0], currentDay1pt2DataIndex))
		currentDay1pt2DataIndex++
	},
}

var calculateAllDay1pt2ClickHandler = jslayer.EventListener{
	Selector:  jslayer.Id(constants.Advent2024Day1Part2CalculateAllID),
	EventType: "click",
	Listener: func(this js.Value, args []js.Value) {
		allTotal := advent2024.Day12CalculateTotal(lookupTable.Value, day1pt2Data.Value[0])	
		fmt.Println("calculateAllDay1pt2ClickHandler", allTotal)
		totalDay1pt2.Set(allTotal)
	},
}

func AdventOfCodeDay1Part2Handler(this js.Value, args []js.Value) any {
	generateLookupTableHandler.Add()
	nextDay1pt2ClickHandler.Add()
	calculateAllDay1pt2ClickHandler.Add()
	day1pt2Data.Set(getJsonData2())
	return nil
}