package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func sendHtml(w http.ResponseWriter, filePath string) {
	wd, err := os.Getwd()

	if err != nil {
		http.Error(w, "Server Side Error", http.StatusInternalServerError)
		return
	}

	html, _ := os.ReadFile(fmt.Sprintf("%s%s", wd, filePath))

	w.Header().Add("Content-Type", "text/html")
	w.Write(html)
}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileHandler := http.StripPrefix("/assets", http.FileServer(http.Dir("/Users/admin/Desktop/inertia-with-go/dist/assets")))

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Header.Get("X-Inertia") == "" && !strings.Contains(r.URL.Path, "assets") {
				sendHtml(w, "/dist/index.html")
				return
			}

			fmt.Printf("%+v\n", r.URL.Path)

			next.ServeHTTP(w, r)
		})
	})

	r.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n***\ni was being called\n***\n")
		fileHandler.ServeHTTP(w, r)
	})

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
