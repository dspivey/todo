package home

import (
	"net/http"
	"todo/router"
	"todo/view"
)

// Load routes for this controller
func Load() {
	router.Get("/home", index)
}

// Index displays the Home page
func index(rw http.ResponseWriter, req *http.Request) {
	template := view.ParseTemplates("home")
	if template == nil {
		rw.Write([]byte("Template 'home.html' not found!"))
	} else {
		template.Execute(rw, nil)
	}
}
