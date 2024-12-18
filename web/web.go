//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/dudubtw/giga-algorithms/controllers"
	jslayer "github.com/dudubtw/giga-algorithms/js-layer"

	advent2024day2 "github.com/dudubtw/giga-algorithms/advent-2024/day-2"
	day3 "github.com/dudubtw/giga-algorithms/advent-2024/day-3"
	day4 "github.com/dudubtw/giga-algorithms/advent-2024/day-4"
)

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("main")

	js.Global().Set("AdventOfCodeDay1Part2Handler", js.FuncOf(controllers.AdventOfCodeDay1Part2Handler))
	js.Global().Set("AdventOfCodeDay1Handler", js.FuncOf(controllers.AdventOfCodeDay1Handler))

	advent2024day2.WebHandlers()
	day3.WebHandlers()
	day4.WebHandlers()

	if !jslayer.IsNil(js.Global().Get("start")) {
		js.Global().Call("start")
	}

	<-c
}
