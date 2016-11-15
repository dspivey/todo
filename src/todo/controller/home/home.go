package home

import (
	"encoding/json"
	"net/http"
	"strconv"
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
		if err != nil {
			log.Danger(err, "Cannot get user from session")
		}

		taskValue := req.PostFormValue("task")

		priorityId, err := strconv.ParseInt(req.PostFormValue("priority"), 10, 64)
		if err != nil {
			log.Danger(err, "Could not parse Priority Id from form request")
		}

		dueAt, err := time.Parse("01/02/2006", req.PostFormValue("due"))
		if err != nil {
			log.Danger(err, "Could not parse Due Date from form request")
		}

		task, err := user.CreateTask(taskValue, int(priorityId), 0, dueAt)
		if err != nil {
			log.Danger(err, "Cannot create task.")
			jsonOutput, _ := json.Marshal(err)
			rw.Write(jsonOutput)
		}

		log.Info(task)

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
