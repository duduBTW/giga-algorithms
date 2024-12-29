//go:build js && wasm
// +build js,wasm

package day3

import (
	"errors"
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/a-h/templ"
	"github.com/dudubtw/giga-algorithms/controllers"
	jslayer "github.com/dudubtw/giga-algorithms/js-layer"
)

var total controllers.StateProps[int]
var expressions controllers.StateProps[[]ExpressionFoundPosition]
var expressionMetadata controllers.StateProps[ExpressionFoundPosition]
var hilightClickHandler jslayer.EventListener
var expressionMouseEnterHandler jslayer.EventListener
var expressionMouseLeaveHandler jslayer.EventListener
var expressionFocusHandler jslayer.EventListener
var expressionBlurHandler jslayer.EventListener
var editorChangeHandler jslayer.EventListener

func getPt1Data() string {
	data, err := jslayer.GetStringAttr(jslayer.Id(IdsData), "data")
	if err != nil {
		fmt.Println("Error getting getPt1Data data day3", err)
		return ""
	}

	return data
}

func getExpressionFromEvent(this js.Value, expressions []ExpressionFoundPosition) (ExpressionFoundPosition, error) {
	attrValue := this.Get("dataset").Get("index")
	if jslayer.IsNil(attrValue) {
		return ExpressionFoundPosition{}, errors.New("Index attr not found")
	}

	index, err := strconv.Atoi(attrValue.String())
	if err != nil {
		return ExpressionFoundPosition{}, err
	}

	return expressions[index], nil
}

func setup() {
	code := getPt1Data()

	expressions = controllers.StateProps[[]ExpressionFoundPosition]{
		Value:  []ExpressionFoundPosition{},
		Target: jslayer.Id(CodeContent),
		RenderComponent: func(value []ExpressionFoundPosition) templ.Component {
			return Part1CodeContent(code, value)
		},
		OnMounted: func(value []ExpressionFoundPosition) {
			jslayer.AddEvents([]jslayer.EventListener{
				expressionMouseEnterHandler,
				expressionMouseLeaveHandler,
				expressionFocusHandler,
				expressionBlurHandler,
				editorChangeHandler,
			})

			elements, err := jslayer.QuerySelectorAll(jslayer.Id(IdsExpression))
			if err != nil {
				return
			}

			elements[0].Call("focus")
		},
	}
	total = controllers.StateProps[int]{
		Value:  0,
		Target: jslayer.Id(IdsTotal),
		RenderComponent: func(value int) templ.Component {
			return Total(value)
		},
	}
	expressionMetadata = controllers.StateProps[ExpressionFoundPosition]{
		Value:  ExpressionFoundPosition{},
		Target: jslayer.Id(IdsExpressionMetadata),
		RenderComponent: func(value ExpressionFoundPosition) templ.Component {
			return ExpressionMetadata(value)
		},
	}

	hilightClickHandler = jslayer.EventListener{
		Selector:  jslayer.Id(IdHilight),
		EventType: "click",
		Listener: func(this js.Value, args []js.Value) {
			foundExpressions := ExpressionFinder(code, []Expression{
				{
					Name: "mul",
					ValidateArgChar: func(r rune) bool {
						_, err := strconv.Atoi(string(r))
						return err == nil
					},
				},
			})
			if len(foundExpressions) == 0 {
				return
			}

			expressions.Set(foundExpressions)

			newTotal := 0
			for _, expression := range foundExpressions {
				if len(expression.Args) < 2 {
					continue
				}

				newTotal += MultiplyArgs(expression.Args)
			}
			total.Set(newTotal)
		},
	}

	var focusExpressionMetadata = func(this js.Value, args []js.Value) {
		expression, err := getExpressionFromEvent(this, expressions.Value)
		if err != nil {
			fmt.Println(err)
			return
		}

		expressionMetadata.Set(expression)
	}
	var blurExpressionMetadata = func(this js.Value, args []js.Value) {
		expressionMetadata.Set(ExpressionFoundPosition{})
	}

	expressionMouseEnterHandler = jslayer.EventListener{
		Selector:  jslayer.Id(IdsExpression),
		EventType: "mouseenter",
		Listener:  focusExpressionMetadata,
	}
	expressionFocusHandler = jslayer.EventListener{
		Selector:  jslayer.Id(IdsExpression),
		EventType: "focus",
		Listener:  focusExpressionMetadata,
	}

	expressionMouseLeaveHandler = jslayer.EventListener{
		Selector:  jslayer.Id(IdsExpression),
		EventType: "mouseleave",
		Listener:  blurExpressionMetadata,
	}
	expressionBlurHandler = jslayer.EventListener{
		Selector:  jslayer.Id(IdsExpression),
		EventType: "blur",
		Listener:  blurExpressionMetadata,
	}

	editorChangeHandler = jslayer.EventListener{
		Selector:  jslayer.Id(CodeContent),
		EventType: "input",
		Listener: func(this js.Value, args []js.Value) {
			fmt.Println(this.Get("innerText").String())
			code = this.Get("innerText").String()
		},
	}
}

func WebHandlers() {
	js.Global().Set(Start, js.FuncOf(func(this js.Value, args []js.Value) any {
		setup()

		hilightClickHandler.Add()
		editorChangeHandler.Add()

		return nil
	}))
}
