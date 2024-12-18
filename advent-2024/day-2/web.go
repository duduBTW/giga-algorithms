//go:build js && wasm
// +build js,wasm

package advent2024day2

import (
	"fmt"
	"syscall/js"

	"github.com/a-h/templ"
	"github.com/dudubtw/giga-algorithms/controllers"
	jslayer "github.com/dudubtw/giga-algorithms/js-layer"
)

func getPt1Data() [][]int {
	newData, err := jslayer.GetJsonData[[][]int](IdsData)
	if err != nil {
		fmt.Println("Error getting getPt1Data data", err)
		return nil
	}

	return newData
}

var selectedFilterOption controllers.StateProps[string]
var radioPart1ClickHandler jslayer.EventListener
var calculateAllPart1Handler jslayer.EventListener
var pt1Data controllers.StateProps[Part1ListProps]
var total controllers.StateProps[int]

func setup() {
	// State
	total = controllers.StateProps[int]{
		Value:  0,
		Target: jslayer.Id(IdsPart1TotalContainer),
		RenderComponent: func(value int) templ.Component {
			return Part1Total(value)
		},
		OnMounted: func(value int) {
			calculateAllPart1Handler.Remove()
		},
	}

	selectedFilterOption = controllers.StateProps[string]{
		Value:  "",
		Target: jslayer.Id(IdsPt1RadioContainer),
		RenderComponent: func(selected string) templ.Component {
			options := getfilterRadioOptions()
			options.SelectedOption = selected

			if selected == "" {
				options.Disabled = true
			}

			return Part1SortRadio(options)
		},
		OnMounted: func(value string) {
			radioPart1ClickHandler.Remove()
			radioPart1ClickHandler.Add()
			js.Global().Get("lucide").Call("createIcons")
		},
	}

	pt1Data = controllers.StateProps[Part1ListProps]{
		Value: Part1ListProps{
			Reports:     getPt1Data(),
			UnsafeIndex: []int{},
		},
		Target: jslayer.Id(IdsPt1List),
		RenderComponent: func(props Part1ListProps) templ.Component {
			return Part1List(SortReportsPropsByType(props, selectedFilterOption.Value))
		},
	}

	// Events
	radioPart1ClickHandler = jslayer.EventListener{
		Selector:  jslayer.Id(IdsPt1Radio),
		EventType: "click",
		Listener: func(this js.Value, args []js.Value) {
			if total.Value == 0 {
				return
			}

			attrValue := this.Get("dataset").Get("value")
			if jslayer.IsNil(attrValue) {
				fmt.Println("Radio option does not have a value!")
				return
			}

			selectedFilterOption.Set(attrValue.String())
			pt1Data.Set(pt1Data.Value)
		},
	}

	calculateAllPart1Handler = jslayer.EventListener{
		Selector:  jslayer.Id(IdsCalculateAll),
		EventType: "click",
		Listener: func(this js.Value, args []js.Value) {
			safeReports := SolvePart1(pt1Data.Value.Reports)
			total.Set(safeReports.SafeSize)
			pt1Data.Set(Part1ListProps{
				Reports:     pt1Data.Value.Reports,
				UnsafeIndex: safeReports.UnsafeIndex,
			})
			selectedFilterOption.Set(FilterOptionAll)
		},
	}
}

func WebHandlers() {
	js.Global().Set("AdventOfCodeDay2Part1Handler", js.FuncOf(func(this js.Value, args []js.Value) any {
		setup()

		calculateAllPart1Handler.Add()
		radioPart1ClickHandler.Add()
		selectedFilterOption.Set("")
		return nil
	}))
}
