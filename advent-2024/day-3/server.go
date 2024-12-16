package day3

import (
	"net/http"

	"github.com/dudubtw/giga-algorithms/components"
)

func Part1ServerHandler() {
	http.HandleFunc("/advent-2024/day-3-part-1", func(w http.ResponseWriter, r *http.Request) {
		content, err := ReadInput("advent-2024/day-3/data-part-1.txt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		component := components.App(Part1Component(Part1ComponentProps{
			Code: content,
		}))
		component.Render(r.Context(), w)
	})
}

func ServerHandlers() {
	Part1ServerHandler()
}
