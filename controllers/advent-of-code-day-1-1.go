//go:build js && wasm
// +build js,wasm

package controllers

import (
	"fmt"
	"syscall/js"

	"context"
	"strings"

	"github.com/a-h/templ"
	advent2024 "github.com/dudubtw/giga-algorithms/advent-2024"
	"github.com/dudubtw/giga-algorithms/components"
	"github.com/dudubtw/giga-algorithms/constants"
	jslayer "github.com/dudubtw/giga-algorithms/js-layer"
)

type StateProps[T any] struct {
	Value           T
	Target          string
	RenderComponent func(value T) templ.Component
	OnMounted       func(value T)
}

func Render(component templ.Component, target string) error {
	componentHTML := new(strings.Builder)
	component.Render(context.Background(), componentHTML)
	return jslayer.ReplaceWithHTML(target, componentHTML.String())
}

func (state *StateProps[T]) Set(value T) {
	state.Value = value
	err := Render(state.RenderComponent(value), state.Target)
	if err != nil {
		fmt.Println("Error setting inner html: ", err)
	}

	if state.OnMounted != nil {
		state.OnMounted(value)
	}
}

func GetJsonData() [][]int {
	data, err := jslayer.GetJsonData[[][]int](constants.Advent2024Day1DataID)
	if err != nil {
		fmt.Println("Error getting json data", err)
		return nil
	}
	return data
}

var currentListIndex = 0
var isSorted = StateProps[bool]{
	Value:  false,
	Target: jslayer.Id(constants.Advent2024Day1SortingContainerID),
	RenderComponent: func(value bool) templ.Component {
		return components.Sorting(value)
	},
}
var data = StateProps[[][]int]{
	Value:  GetJsonData(),
	Target: jslayer.Id(constants.Advent2024Day1ListID),
	RenderComponent: func(value [][]int) templ.Component {
		return components.Day1List(value)
	},
}
var total = StateProps[int]{
	Value:  0,
	Target: jslayer.Id(constants.Advent2024Day1TotalID),
	RenderComponent: func(value int) templ.Component {
		return components.TotalValue(value)
	},
}

var calculateClickHandler = jslayer.EventListener{
	Selector:  jslayer.Id(constants.Advent2024Day1CalculateAllID),
	EventType: "click",
	Listener: func(this js.Value, args []js.Value) {
		newTotal, newListWithDiffs := advent2024.CalculateAll(data.Value)
		total.Set(newTotal)
		data.Set(newListWithDiffs)
	},
}

var nextClickHandler = jslayer.EventListener{
	Selector:  jslayer.Id(constants.Advent2024Day1NextID),
	EventType: "click",
	Listener: func(this js.Value, args []js.Value) {
		diff := advent2024.Day11Line(data.Value, currentListIndex)
		newData := data.Value
		newData[2][currentListIndex] = diff
		data.Set(newData)
		total.Set(total.Value + diff)
		currentListIndex++
	},
}

var clearAllClickHandler = jslayer.EventListener{
	Selector:  jslayer.Id(constants.Advent2024Day1ClearAllID),
	EventType: "click",
	Listener: func(this js.Value, args []js.Value) {
		data.Set(make([][]int, 3))
	},
}

var sortClickHandler = jslayer.EventListener{
	Selector:  jslayer.Id(constants.Advent2024Day1SortID),
	EventType: "click",
	Listener: func(this js.Value, args []js.Value) {
		fmt.Println("sortClickHandler")
		data.Set(advent2024.SortLines(data.Value))
		isSorted.Set(true)
	},
}

func AdventOfCodeDay1Handler(this js.Value, args []js.Value) any {
	calculateClickHandler.Add()
	sortClickHandler.Add()
	nextClickHandler.Add()
	clearAllClickHandler.Add()
	return nil
}
