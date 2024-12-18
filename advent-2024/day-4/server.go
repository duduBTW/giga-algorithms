package day4

import (
	"net/http"

	"github.com/dudubtw/giga-algorithms/components"
)

func Part1ServerHandler() {
	http.HandleFunc("/advent-2024/day-4-part-1", func(w http.ResponseWriter, r *http.Request) {
		matrix, err := ReadInput("D:/Peronal/giga-algorithms/advent-2024/day-4/data-part-1.txt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		component := components.App(Part1Component(Part1ComponentProps{
			Matrix: matrix,
		}))
		component.Render(r.Context(), w)
	})
}

func ServerHandlers() {
	Part1ServerHandler()
}
