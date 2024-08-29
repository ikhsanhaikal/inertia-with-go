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

	r.Get("/page1", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Component string      `json:"component"`
			Props     interface{} `json:"props"`
			Url       string      `json:"url"`
			Version   string      `json:"version"`
		}{
			Component: "Page1",
			Props: struct {
				Data Todo `json:"data"`
			}{
				Data: Todo{
					Id:        1,
					Task:      "take out the trash",
					Completed: true,
				},
			},
			Url:     "/page1",
			Version: "",
		}

		b, err := json.Marshal(data)

		fmt.Printf("page object: %s\n", string(b))

		if err != nil {
			http.Error(w, "Server Side Error", http.StatusInternalServerError)
			return
		}

		sendJson(w, b)
	})

	http.ListenAndServe("localhost:4040", r)
}

type Todo struct {
	Id        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func sendJson(w http.ResponseWriter, data []byte) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Vary", "X-Inertia")
	w.Header().Set("X-Inertia", "true")
	fmt.Fprint(w, string(data))
}
