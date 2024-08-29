package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileHandler := http.StripPrefix("/assets", http.FileServer(http.Dir("/Users/admin/Desktop/inertia-with-go/dist/assets")))

	r.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n***\ni was being called\n***\n")
		fileHandler.ServeHTTP(w, r)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageObject{
			Component: "Index",
			Props:     "{}",
			Url:       "/",
			Version:   "",
		}

		inertiaResponse(w, r, data)
	})

	r.Get("/page1", func(w http.ResponseWriter, r *http.Request) {
		data := PageObject{
			Component: "Page1",
			Props: struct {
				Data Todo `json:"data"`
			}{
				Data: Todo{
					Id:        11,
					Task:      "take out the trash",
					Completed: false,
				}},
			Url:     "/page1",
			Version: "",
		}

		inertiaResponse(w, r, data)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		inertiaResponse(w, r, PageObject{
			Component: "Notfound",
			Props:     "{}",
			Url:       "/404",
			Version:   "",
		})
	})

	http.ListenAndServe("localhost:4040", r)
}

func inertiaResponse(w http.ResponseWriter, r *http.Request, pageObject PageObject) {
	if r.Header.Get("X-Inertia") == "" {
		wd, err := os.Getwd()

		if err != nil {
			http.Error(w, "Server Side Error", http.StatusInternalServerError)
			return
		}

		t1 := template.New("index")
		t1, err = t1.ParseFiles(fmt.Sprintf("%s/dist/index.html", wd))

		if err != nil {
			http.Error(w, "Server Side Error", http.StatusInternalServerError)
			return
		}

		props, err := json.Marshal(pageObject.Props)

		if err != nil {
			http.Error(w, "Server Side Error", http.StatusInternalServerError)
			return
		}

		t1.ExecuteTemplate(w, "index.html", PageObject{
			Component: pageObject.Component,
			Props:     string(props),
			Url:       pageObject.Url,
			Version:   pageObject.Version,
		})

		return
	}

	b, err := json.Marshal(pageObject)

	if err != nil {
		http.Error(w, "Server Side Error", http.StatusInternalServerError)
		return
	}
	sendJson(w, b)
}

type PageObject struct {
	Component string      `json:"component"`
	Props     interface{} `json:"props"`
	Url       string      `json:"url"`
	Version   string      `json:"version"`
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
