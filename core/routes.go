package core

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func MainRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
}
