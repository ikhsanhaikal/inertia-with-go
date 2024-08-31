package handler

import "net/http"

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	data := PageObject{
		Component: "Index",
		Props:     "{}",
		Url:       "/",
		Version:   "",
	}

	app.inertiaResponse(w, r, data)
}
