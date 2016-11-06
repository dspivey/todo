package home

import (
	"net/http"
	"todo/log"
	"todo/model"
	"todo/router"
	"todo/view"
)

// Load routes for this controller
func Load() {
	router.Get("/", index)
	router.Get("/home", index)
}

// Index displays the Home page
func index(rw http.ResponseWriter, req *http.Request) {
	s, err := model.IsAuthenticated(rw, req)
	if err != nil {
		http.Redirect(rw, req, "/login", 302)
	} else {
		user, err := s.User()
		if err != nil {
			log.Danger(err, "Cannot get user from session")
		}

		views := []string{"shared/_layout", "home/home"}
		view.RenderHTML(rw, user, views...)
	}
}
