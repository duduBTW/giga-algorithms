package day5

import (
	"net/http"

	"github.com/dudubtw/giga-algorithms/components"
)

func ServerHandlers() {
	http.HandleFunc(Route, func(w http.ResponseWriter, r *http.Request) {
		manual, err := ReadInput()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		component := components.App(Component(ComponentProps{
			Manual: manual,
		}))
		component.Render(r.Context(), w)
	})
}
