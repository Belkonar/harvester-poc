package main

import (
	"fmt"
	"harvester_api/core"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	core.MainRoutes(r)

	fmt.Println("Starting on port :3000")
	http.ListenAndServe(":3000", r)
}
