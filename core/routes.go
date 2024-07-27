package core

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func MainRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := WriteKV("/e/f", "lolgd")

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("good"))
	})
}
