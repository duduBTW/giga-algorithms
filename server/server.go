package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/dudubtw/giga-algorithms/components"
)

func main() {
	// PUBLIC SERVER
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public", fileServer))

	// HANDLE ROOT
	http.Handle("/", templ.Handler(components.App(components.Home())))

	// SERVE APP
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
