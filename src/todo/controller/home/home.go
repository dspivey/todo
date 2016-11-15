package home

import (
	"net/http"
	"todo/log"
	"todo/model"
	"todo/router"
	"todo/view"
	"todo/viewmodel"
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

		tasks, err := user.Tasks()
		if err != nil {
			log.Danger(err, "Cannot get tasks for user")
		}

		priorities, err := model.Priorities()
		if err != nil {
			log.Danger(err, "Cannot get priorities")
		}

		tags, err := model.Tags()
		if err != nil {
			log.Danger(err, "Cannot get priorities")
		}

		vm := viewmodel.HomeViewModel{
			Title:      "Doozer Checklist",
			User:       user,
			Tasks:      tasks,
			Priorities: priorities,
			Tags:       tags,
		}

		views := []string{"shared/_layout", "shared/_navigation", "home/home"}
		view.RenderHTML(rw, vm, "view", views...)
	}
}