package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func (app *Application) inertiaResponse(w http.ResponseWriter, r *http.Request, pageObject PageObject) {
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

func sendJson(w http.ResponseWriter, data []byte) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Vary", "X-Inertia")
	w.Header().Set("X-Inertia", "true")
	fmt.Fprint(w, string(data))
}

type PageObject struct {
	Component string      `json:"component"`
	Props     interface{} `json:"props"`
	Url       string      `json:"url"`
	Version   string      `json:"version"`
}

type Todo struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Task      string `json:"title"`
	Completed bool   `json:"completed"`
}
