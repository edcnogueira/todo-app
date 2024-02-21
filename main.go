package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("John")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return
	}
}
