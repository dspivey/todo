package tasks

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
	router.Get("/tasks", tasks)
	router.Post("/tasks/create", createTask)
	router.Post("/tasks/test/tags", testTags)
}

func tasks(rw http.ResponseWriter, req *http.Request) {
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

		isJsonReq, _ := strconv.ParseBool(req.URL.Query().Get("json"))
		if isJsonReq == true {
			jsonOutput, _ := json.Marshal(&tasks)

        	rw.Header().Set("content-type", "application/json")
        	rw.Write(jsonOutput)
		} else {
			vm := viewmodel.TasksViewModel {
				User: user,
				Tasks: tasks,
			}

			views := []string{"shared/_tasks"}
			view.RenderHTML(rw, vm, "tasks", views...)
		}
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

		// get user from session
		user, err := s.User()
		if err != nil {
			log.Danger(err, "Cannot get user from session")
		}

		// grab values from post
		taskValue := req.PostFormValue("task")

		priorityId, err := strconv.ParseInt(req.PostFormValue("priority"), 10, 64)
		if err != nil {
			log.Danger(err, "Could not parse Priority Id from form request")
		}

		dueAt, err := time.Parse("01/02/2006 3:04 PM", req.PostFormValue("due"))
		if err != nil {
			log.Danger(err, "Could not parse Due Date from form request")
		}

		// create a new task
		task, err := user.CreateTask(taskValue, int(priorityId), 0, dueAt)
		if err != nil {
			log.Danger(err, "Cannot create task.")
			jsonOutput, _ := json.Marshal(err)
			rw.Write(jsonOutput)
		}

		// insert any new tags 
		tagsStr := req.PostFormValue("tags")
		log.Info(tagsStr)

		var tagIds []string
		err = json.Unmarshal([]byte(tagsStr), &tagIds)

		_, err = task.AddTags(tagIds)
		if err != nil {
			log.Danger(err, "Cannot create task.")
			jsonOutput, _ := json.Marshal(err)
			rw.Write(jsonOutput)
		}

		// return the newly created task
		jsonOutput, err := json.Marshal(&task)
		if err != nil {
			log.Danger(err, "Cannot marshal task to JSON.")
			jsonOutput, _ := json.Marshal(err)

			rw.Write(jsonOutput)
		}

		rw.Write(jsonOutput)
	}
}

func testTags(rw http.ResponseWriter, req *http.Request){
	err := req.ParseForm()
	if err != nil {
		log.Danger(err, "Cannot parse form")
		jsonOutput, _ := json.Marshal(err)

		rw.Write(jsonOutput)
	}

	tagsStr := req.PostFormValue("tags")
	log.Info(tagsStr)

	var tags []string
	err = json.Unmarshal([]byte(tagsStr), &tags)

	log.Info(tags)

}
