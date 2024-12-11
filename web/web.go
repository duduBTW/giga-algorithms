//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/dudubtw/giga-algorithms/controllers"

	advent2024day2 "github.com/dudubtw/giga-algorithms/advent-2024/day-2"
)

func main() {
	fmt.Println("Test a")
	c := make(chan struct{}, 0)

	fmt.Println("main")

	js.Global().Set("AdventOfCodeDay1Part2Handler", js.FuncOf(controllers.AdventOfCodeDay1Part2Handler))
	js.Global().Set("AdventOfCodeDay1Handler", js.FuncOf(controllers.AdventOfCodeDay1Handler))

	advent2024day2.WebHandlers()

	js.Global().Call("start")

	<-c
}
