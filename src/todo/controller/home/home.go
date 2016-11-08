package home

import (
	"encoding/json"
	"net/http"
	"time"
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
	router.Post("/tasks/create", createTask)
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

		vm := viewmodel.HomeViewModel{
			Title: "Doozer Checklist",
			User:  user,
			Tasks: tasks,
		}

		views := []string{"shared/_layout", "shared/_navigation", "home/home"}
		view.RenderHTML(rw, vm, views...)
	}
}

// POST /createTask
// Create the user account
func createTask(rw http.ResponseWriter, req *http.Request) {
	s, err := model.IsAuthenticated(rw, req)
	if err != nil {
		http.Redirect(rw, req, "/login", 302)
	} else {
		rw.Header().Set("content-type", "application/json")
		err := req.ParseForm()
		if err != nil {
			log.Danger(err, "Cannot parse form")
			jsonOutput, _ := json.Marshal(err)
			rw.Write(jsonOutput)
		}

		user, err := s.User()
		task := model.Task{
			Value:      req.PostFormValue("task"),
			User:       &user,
			StatusId:   4,
			PriorityId: 5,
		}

		task, err = user.CreateTask(task.Value, task.PriorityId, task.StatusId, time.Now())
		if err != nil {
			log.Danger(err, "Cannot create task.")
			jsonOutput, _ := json.Marshal(err)
			rw.Write(jsonOutput)
		}

		log.Info(task.Value)

		jsonOutput, err := json.Marshal(&task)
		if err != nil {
			log.Danger(err, "Cannot marshal task to JSON.")
			jsonOutput, _ := json.Marshal(err)
			rw.Write(jsonOutput)
		}

		rw.Header().Set("content-type", "application/json")
		rw.Write(jsonOutput)
	}
}
