package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		msg := struct {
			Message    string `json:"message"`
			StatusCode int    `json:"statusCode"`
		}{
			Message:    "ok",
			StatusCode: 200,
		}

		data, err := json.Marshal(msg)

		if err != nil {
			http.Error(w, "meh", 200)
			return
		}

		respJson(w, data)
	})

	http.ListenAndServe("localhost:4040", r)
}

func respJson(w http.ResponseWriter, data []byte) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(data))
}
