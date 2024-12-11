package advent2024day2

import (
	"net/http"

	"github.com/dudubtw/giga-algorithms/components"
)

func Part1ServerHandler() {
	http.HandleFunc("/advent-2024/day-2-part-1", func(w http.ResponseWriter, r *http.Request) {
		reports, err := ReadDay2Input("advent-2024/day-2/data-part-1.txt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		component := components.App(Day2Part1Component(Day2Part1Props{
			Reports: reports,
		}))
		component.Render(r.Context(), w)
	})
}

func ServerHandlers() {
	Part1ServerHandler()
}
