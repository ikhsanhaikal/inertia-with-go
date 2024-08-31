package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (app *Application) ApiPage2Get(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("query: %+v\n", r.URL.Query())
	queries := r.URL.Query()
	limit := 5
	page := 1

	if queries.Has("limit") {
		l, _ := strconv.Atoi(queries.Get("limit"))
		limit = l
	}

	if queries.Has("page") {
		p, _ := strconv.Atoi(queries.Get("page"))
		page = p
	}

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos?_page=%d&_limit=%d", page, limit)

	fmt.Printf("url: %s\n", url)

	response, err := http.Get(url)

	if err != nil {
		http.Error(w, "Server Side Error", http.StatusInternalServerError)
		return
	}

	b, err := io.ReadAll(response.Body)

	if err != nil {
		http.Error(w, "Server Side Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(b))
}

func (app *Application) Page2Get(w http.ResponseWriter, r *http.Request) {

	props := struct {
		InMemoryData []Todo `json:"inMemoryData"`
	}{
		InMemoryData: app.InMemoryData,
	}

	b, _ := json.Marshal(props)

	fmt.Printf("json: %s\n", string(b))

	app.inertiaResponse(w, r, PageObject{
		Component: "Page2Index",
		Props:     props,
		Version:   "",
		Url:       "/page2",
	})
}

func (app *Application) Page2New(w http.ResponseWriter, r *http.Request) {
	app.inertiaResponse(w, r, PageObject{
		Component: "Page2New",
		Props:     "{}",
		Version:   "",
		Url:       "/page2/new",
	})
}

func (app *Application) Page2Create(w http.ResponseWriter, r *http.Request) {
	content, _ := io.ReadAll(r.Body)
	fmt.Printf("content: %s", string(content))

	t := Todo{}

	if err := json.Unmarshal(content, &t); err != nil {
		http.Error(w, "Server Side Error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("t: %+v\n", t)

	app.InMemoryData = append(app.InMemoryData, t)

	props := struct {
		InMemoryData []Todo `json:"inMemoryData"`
	}{
		InMemoryData: app.InMemoryData,
	}

	app.inertiaResponse(w, r, PageObject{
		Component: "Page2Index",
		Props:     props,
		Version:   "",
		Url:       "/page2",
	})
}
