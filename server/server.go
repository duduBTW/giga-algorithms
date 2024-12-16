package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	advent2024 "github.com/dudubtw/giga-algorithms/advent-2024"
	advent2024dat2 "github.com/dudubtw/giga-algorithms/advent-2024/day-2"
	day3 "github.com/dudubtw/giga-algorithms/advent-2024/day-3"
	"github.com/dudubtw/giga-algorithms/components"
)

func main() {
	// Create a new server
	srv := &http.Server{
		Addr: ":9081",
	}

	// PUBLIC SERVER
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public", fileServer))

	// HANDLE ROOT
	http.Handle("/", templ.Handler(components.App(components.Home())))
	http.HandleFunc("/advent-2024/day-1", advent2024Day1Handler)
	http.HandleFunc("/advent-2024/day-1-part-2", advent2024Day1Part2Handler)

	advent2024dat2.ServerHandlers()
	day3.ServerHandlers()

	// Start server in a goroutine
	fmt.Println("Server is running on port 8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func advent2024Day1Handler(w http.ResponseWriter, r *http.Request) {
	day1Data, err := advent2024.ReadDay1Input("day-1-1.txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := components.App(components.Advent2024Day1(components.Advent2024Day1Props{
		Data: day1Data,
	}))
	component.Render(r.Context(), w)
}

func advent2024Day1Part2Handler(w http.ResponseWriter, r *http.Request) {
	day1Data, err := advent2024.ReadDay1Input("day-1-2.txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := components.App(components.Advent2024Day1Part2(components.Advent2024Day1Part2Props{
		Data: day1Data,
	}))
	component.Render(r.Context(), w)
}
