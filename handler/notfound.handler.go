package handler

import "net/http"

func (app *Application) NotFound(w http.ResponseWriter, r *http.Request) {
	app.inertiaResponse(w, r, PageObject{
		Component: "Notfound",
		Props:     "{}",
		Url:       "/404",
		Version:   "",
	})
}
