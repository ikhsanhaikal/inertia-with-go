package main

import (
	"fmt"
	"net/http"

	"com.ikhsanhaikal.inertiawithgo/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	app := handler.Application{
		InMemoryData: []handler.Todo{
			{Id: 14, UserId: 4, Task: "task: learn how to networking", Completed: false},
			{Id: 51, UserId: 11, Task: "task: learn to hold breath for 6 min", Completed: false},
		},
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileHandler := http.StripPrefix("/assets", http.FileServer(http.Dir("/Users/admin/Desktop/inertia-with-go/dist/assets")))

	r.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n***\ni was being called\n***\n")
		fileHandler.ServeHTTP(w, r)
	})

	r.Get("/", app.Home)
	r.Get("/page1", app.Page1)

	r.Get("/page2", app.Page2Get)
	r.Get("/api/page2", app.ApiPage2Get)
	r.Get("/page2/new", app.Page2New)
	r.Post("/page2", app.Page2Create)

	r.NotFound(app.NotFound)

	http.ListenAndServe("localhost:4040", r)
}
