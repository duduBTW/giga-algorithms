//go:build js && wasm
// +build js,wasm

package day4

import (
	"fmt"
	"sort"
	"strconv"
	"syscall/js"

	jslayer "github.com/dudubtw/giga-algorithms/js-layer"
)

// Function to compare positions by X, then by Y
func (p Position) Less(p2 Position) bool {
	if p.X == p2.X {
		return p.Y < p2.Y
	}
	return p.X < p2.X
}

func convexHull(positions []Position) []Position {
	// Sort the positions lexicographically (by X, then Y)
	sort.Slice(positions, func(i, j int) bool {
		return positions[i].Less(positions[j])
	})

	// Build lower hull
	var lower []Position
	for _, p := range positions {
		for len(lower) >= 2 && crossProduct(lower[len(lower)-2], lower[len(lower)-1], p) <= 0 {
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, p)
	}

	// Build upper hull
	var upper []Position
	for i := len(positions) - 1; i >= 0; i-- {
		for len(upper) >= 2 && crossProduct(upper[len(upper)-2], upper[len(upper)-1], positions[i]) <= 0 {
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, positions[i])
	}

	// Remove the last point of each half because it is repeated at the beginning of the other half
	return append(lower[:len(lower)-1], upper[:len(upper)-1]...)
}

// Cross product to determine if positions make a counter-clockwise turn
func crossProduct(p1, p2, p3 Position) int {
	return (p2.X-p1.X)*(p3.Y-p1.Y) - (p2.Y-p1.Y)*(p3.X-p1.X)
}

var formSubmitHandler jslayer.EventListener

func getMatrix() [][]rune {
	matrix, err := jslayer.GetJsonData[[][]rune](IdData)
	if err != nil {
		fmt.Println("Error getting data day 4", err)
		return nil
	}

	return matrix
}

func GetCell(position Position) (js.Value, error) {
	return jslayer.QuerySelector(jslayer.Id(IdCell) + `[data-x-index="` + strconv.Itoa(position.X) + `"][data-y-index="` + strconv.Itoa(position.Y) + `"]`)
}

func DrawHighlights(highlights []Highlight) {
	canvas, err := jslayer.QuerySelector(jslayer.Id(IdCanvas))
	if err != nil {
		fmt.Println("Could not draw highlight!")
		return
	}

	context := canvas.Call("getContext", "2d")
	container, err := jslayer.QuerySelector(jslayer.Id(IdMatrix))
	if err != nil {
		fmt.Println("Matrix element not found!")
		return
	}

	width := container.Get("scrollWidth").Int()
	height := container.Get("scrollHeight").Int()
	canvas.Set("width", width)
	canvas.Set("height", height)

	containerBox := container.Call("getBoundingClientRect")

	// Clear canvas
	canvasWidth := canvas.Get("width").Int()
	canvasHeight := canvas.Get("height").Int()
	context.Call("clearRect", 0, 0, canvasWidth, canvasHeight)

	for _, highlight := range highlights {
		startCell, err := GetCell(highlight.Start)
		if err != nil {
			fmt.Println("start cell element not found!")
			continue
		}

		endCell, err := GetCell(highlight.End)
		if err != nil {
			fmt.Println("end cell element not found!")
			continue
		}

		boundingBox1 := startCell.Call("getBoundingClientRect")
		box1ScrollWidth := startCell.Get("scrollWidth").Float()
		boundingBox2 := endCell.Call("getBoundingClientRect")
		box2ScrollWidth := endCell.Get("scrollWidth").Float()

		// Calculate the container's padding
		paddingTop := containerBox.Get("top").Float()
		paddingLeft := containerBox.Get("left").Float()

		js.Global().Call("scrollTo", 0, 0)
		container.Call("scrollTo", 0, 0)

		// Collect positions (corners of the divs, adjusted for padding)
		positions := []Position{
			{X: int(boundingBox1.Get("left").Float() - paddingLeft), Y: int(boundingBox1.Get("top").Float() - paddingTop)},
			{X: int(boundingBox1.Get("left").Float() + box1ScrollWidth - paddingLeft), Y: int(boundingBox1.Get("top").Float() - paddingTop)},
			{X: int(boundingBox1.Get("left").Float() - paddingLeft), Y: int(boundingBox1.Get("top").Float() + boundingBox1.Get("height").Float() - paddingTop)},
			{X: int(boundingBox1.Get("left").Float() + box1ScrollWidth - paddingLeft), Y: int(boundingBox1.Get("top").Float() + boundingBox1.Get("height").Float() - paddingTop)},
			{X: int(boundingBox2.Get("left").Float() - paddingLeft), Y: int(boundingBox2.Get("top").Float() - paddingTop)},
			{X: int(boundingBox2.Get("left").Float() + box2ScrollWidth - paddingLeft), Y: int(boundingBox2.Get("top").Float() - paddingTop)},
			{X: int(boundingBox2.Get("left").Float() - paddingLeft), Y: int(boundingBox2.Get("top").Float() + boundingBox2.Get("height").Float() - paddingTop)},
			{X: int(boundingBox2.Get("left").Float() + box2ScrollWidth - paddingLeft), Y: int(boundingBox2.Get("top").Float() + boundingBox2.Get("height").Float() - paddingTop)},
		}

		hull := convexHull(positions)

		context.Set("strokeStyle", "rgb(59 130 246)")
		context.Set("lineWidth", 1)
		context.Call("beginPath")

		context.Call("moveTo", hull[0].X, hull[0].Y)

		for _, p := range hull[1:] {
			context.Call("lineTo", p.X, p.Y)
		}

		context.Call("closePath")
		context.Call("stroke")
	}
}

func setup() {
	matrix := getMatrix()

	formSubmitHandler = jslayer.EventListener{
		Selector:  jslayer.Id(IdSearchForm),
		EventType: "submit",
		Listener: func(this js.Value, args []js.Value) {
			args[0].Call("preventDefault")

			inputElement, err := jslayer.QuerySelector(jslayer.Id(IdSearchInput))
			if err != nil {
				return
			}

			searchValue := inputElement.Get("value").String()
			go func() {
				highlights := FindWordInstances(searchValue, matrix)
				DrawHighlights(highlights)
			}()
		},
	}
}

func WebHandlers() {
	js.Global().Set(Start, js.FuncOf(func(this js.Value, args []js.Value) any {
		setup()

		formSubmitHandler.Add()

		return nil
	}))
}
