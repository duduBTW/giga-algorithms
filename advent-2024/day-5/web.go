//go:build js && wasm
// +build js,wasm

package day5

import (
	"fmt"
	"syscall/js"

	"github.com/a-h/templ"
	"github.com/dudubtw/giga-algorithms/components"
	"github.com/dudubtw/giga-algorithms/controllers"
	jslayer "github.com/dudubtw/giga-algorithms/js-layer"
)

func getManual() Manual {
	manual, _ := jslayer.GetJsonData[Manual](IdData)
	fmt.Println(manual)
	// TODO - show error page when theres no manual
	return manual
}

var tabsState controllers.StateProps[TabsProps]
var pagesState controllers.StateProps[PagesProps]
var totalState controllers.StateProps[int]
var tabClickHandler jslayer.EventListener
var findClickHandler jslayer.EventListener
var fixClickHandler jslayer.EventListener

func setup() {
	defaultManual := getManual()

	tabsState = controllers.StateProps[TabsProps]{
		Value: TabsProps{
			ActiveTab: "all",
			Manual:    defaultManual,
		},
		Target: jslayer.Id(IdTabs),
		RenderComponent: func(props TabsProps) templ.Component {
			return Tabs(props)
		},
		OnMounted: func(value TabsProps) {
			tabClickHandler.Add()
		},
	}

	pagesState = controllers.StateProps[PagesProps]{
		Value: PagesProps{
			Pages:              defaultManual.Pages,
			InvalidPageIndexes: make([][]int, len(defaultManual.Pages)),
		},
		Target: jslayer.Id(IdPages),
		RenderComponent: func(props PagesProps) templ.Component {
			return PagesComponent(props)
		},
	}

	totalState = controllers.StateProps[int]{
		Value:  0,
		Target: jslayer.Id(IdTotal),
		RenderComponent: func(value int) templ.Component {
			return Total(value)
		},
	}

	tabClickHandler = jslayer.EventListener{
		Selector:  jslayer.Id(IdTabItem),
		EventType: "click",
		Listener: func(this js.Value, args []js.Value) {
			attrValue := this.Get("dataset").Get(components.TabItemValueAttr)
			if jslayer.IsNil(attrValue) {
				fmt.Println("Tab option does not have a value!")
				return
			}

			fmt.Println(defaultManual)
			tabsState.Set(TabsProps{
				ActiveTab: attrValue.String(),
				Manual:    defaultManual,
			})
		},
	}

	findClickHandler = jslayer.EventListener{
		Selector:  jslayer.Id(IdFindButton),
		EventType: "click",
		Listener: func(this js.Value, args []js.Value) {
			newTotal, invalidPageIndexes := ValidateManual(defaultManual)
			totalState.Set(newTotal)
			pagesState.Set(PagesProps{
				Pages:              defaultManual.Pages,
				InvalidPageIndexes: invalidPageIndexes,
			})
		},
	}

	fixClickHandler = jslayer.EventListener{
		Selector:  jslayer.Id(IdFixButton),
		EventType: "click",
		Listener: func(this js.Value, args []js.Value) {
			total, fixedPages, invalidPageIndexes := FixManual(defaultManual)
			totalState.Set(total)
			pagesState.Set(PagesProps{
				Pages:              fixedPages,
				InvalidPageIndexes: invalidPageIndexes,
			})
		},
	}
}

func WebHandlers() {
	js.Global().Set(Start, js.FuncOf(func(this js.Value, args []js.Value) any {
		setup()

		jslayer.AddEvents([]jslayer.EventListener{
			tabClickHandler,
			findClickHandler,
			fixClickHandler,
		})

		return nil
	}))
}
