package handler

import "net/http"

func (app *Application) Page1(w http.ResponseWriter, r *http.Request) {
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

	app.inertiaResponse(w, r, data)
}
