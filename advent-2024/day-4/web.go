//go:build js && wasm
// +build js,wasm

package day4

import (
	"fmt"
	"sort"
	"strconv"
	"syscall/js"
	"time"

	"github.com/a-h/templ"
	"github.com/dudubtw/giga-algorithms/controllers"
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

func getMatrix() [][]rune {
	matrix, err := jslayer.GetJsonData[[][]rune](IdData)
	if err != nil {
		fmt.Println("Error getting data day 4", err)
		return nil
	}

	return matrix
}

var cellCache = make(map[Position]js.Value)

func GetCell(position Position) (js.Value, error) {
	cellFromCache, ok := cellCache[position]
	if ok {
		return cellFromCache, nil
	}

	newCell, err := jslayer.QuerySelector(jslayer.Id(IdCell) + `[data-x-index="` + strconv.Itoa(position.X) + `"][data-y-index="` + strconv.Itoa(position.Y) + `"]`)
	if err != nil {
		return newCell, err
	}

	cellCache[position] = newCell
	return newCell, nil
}

type Box struct {
	TopLeft     Position
	TopRight    Position
	BottomLeft  Position
	BottomRight Position
}

var boxMap = make(map[Position]Box)

func getBoundingBoxPositions(boundingBox js.Value, paddingLeft, paddingTop float64, position Position) Box {
	cachedBox, ok := boxMap[position]
	if ok {
		return cachedBox
	}

	box := Box{
		TopLeft: Position{
			X: int(boundingBox.Get("left").Float() - paddingLeft),
			Y: int(boundingBox.Get("top").Float() - paddingTop),
		},
		TopRight: Position{
			X: int(boundingBox.Get("left").Float() + boundingBox.Get("width").Float() - paddingLeft),
			Y: int(boundingBox.Get("top").Float() - paddingTop),
		},
		BottomLeft: Position{
			X: int(boundingBox.Get("left").Float() - paddingLeft),
			Y: int(boundingBox.Get("top").Float() + boundingBox.Get("height").Float() - paddingTop),
		},
		BottomRight: Position{
			X: int(boundingBox.Get("left").Float() + boundingBox.Get("width").Float() - paddingLeft),
			Y: int(boundingBox.Get("top").Float() + boundingBox.Get("height").Float() - paddingTop),
		},
	}
	boxMap[position] = box
	return box
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

	start := time.Now()

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
		// box1ScrollWidth := startCell.Get("scrollWidth").Float()
		boundingBox2 := endCell.Call("getBoundingClientRect")
		// box2ScrollWidth := endCell.Get("scrollWidth").Float()

		// Calculate the container's padding
		paddingTop := containerBox.Get("top").Float()
		paddingLeft := containerBox.Get("left").Float()

		js.Global().Call("scrollTo", 0, 0)
		container.Call("scrollTo", 0, 0)

		box1 := getBoundingBoxPositions(boundingBox1, paddingLeft, paddingTop, highlight.Start)
		box2 := getBoundingBoxPositions(boundingBox2, paddingLeft, paddingTop, highlight.End)

		positions := []Position{
			box1.TopLeft, box1.TopRight, box1.BottomLeft, box1.BottomRight,
			box2.TopLeft, box2.TopRight, box2.BottomLeft, box2.BottomRight,
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

	elapsed := time.Since(start) // Calculate elapsed time
	fmt.Printf("Function took %s to run\n", elapsed)
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var allDirections = map[string]Direction{
	UpLeft:    {X: Previous, Y: Previous},
	Up:        {X: Neutral, Y: Previous},
	UpRight:   {X: Next, Y: Previous},
	DownLeft:  {X: Previous, Y: Next},
	Down:      {X: Neutral, Y: Next},
	DownRight: {X: Next, Y: Next},
	Left:      {X: Previous, Y: Neutral},
	Right:     {X: Next, Y: Neutral},
}

var angledDirections = map[string]Direction{
	UpLeft:    {X: Previous, Y: Previous},
	UpRight:   {X: Next, Y: Previous},
	DownLeft:  {X: Previous, Y: Next},
	DownRight: {X: Next, Y: Next},
}

var formSubmitHandler jslayer.EventListener
var findXClickHandler jslayer.EventListener
var total controllers.StateProps[int]

func setup() {
	matrix := getMatrix()

	total = controllers.StateProps[int]{
		Value:  0,
		Target: jslayer.Id(IdTotal),
		RenderComponent: func(value int) templ.Component {
			return TotalResults(value)
		},
	}

	var getInputValue = func() (string, error) {
		inputElement, err := jslayer.QuerySelector(jslayer.Id(IdSearchInput))
		if err != nil {
			return "", err
		}
		return inputElement.Get("value").String(), nil
	}

	formSubmitHandler = jslayer.EventListener{
		Selector:  jslayer.Id(IdSearchForm),
		EventType: "submit",
		Listener: func(this js.Value, args []js.Value) {
			args[0].Call("preventDefault")
			searchValue, err := getInputValue()
			if err != nil {
				return
			}

			highlights := FindWordInstances(searchValue, matrix, allDirections)
			total.Set(len(highlights))
			DrawHighlights(highlights)
		},
	}

	findXClickHandler = jslayer.EventListener{
		Selector:  jslayer.Id(IdFindX),
		EventType: "click",
		Listener: func(this js.Value, args []js.Value) {
			searchValue, err := getInputValue()
			if err != nil {
				return
			}

			highlights := FindWordInstances(searchValue, matrix, angledDirections)
			middlePoints := make(map[string]Highlight)
			currentTotal := 0
			validH := []Highlight{}

			for _, highlight := range highlights {
				minY := MaxInt(highlight.Start.Y, highlight.End.Y)
				maxX := MaxInt(highlight.Start.X, highlight.End.X)

				middlePoint := strconv.Itoa(minY-1) + "-" + strconv.Itoa(maxX-1)

				cHighlight, ok := middlePoints[middlePoint]
				if ok {
					currentTotal++
					fmt.Println(highlight, cHighlight)
					validH = append(validH, highlight, cHighlight)
					delete(middlePoints, middlePoint)
					continue
				}

				middlePoints[middlePoint] = highlight
			}

			total.Set(currentTotal)
			DrawHighlights(validH)
		},
	}
}

func WebHandlers() {
	js.Global().Set(Start, js.FuncOf(func(this js.Value, args []js.Value) any {
		setup()

		formSubmitHandler.Add()
		findXClickHandler.Add()

		return nil
	}))
}
