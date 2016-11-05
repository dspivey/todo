package model

import (
	"testing"
	"time"
)

// Delete all threads from database
func setup() (err error){
	statement := `delete from task_tags;
	delete from tags;
	delete from user_tasks;
	delete from tasks;
	delete from sessions;
	delete from users;`

	_, err = Database.Exec(statement)

	return err
}

func Test_CreateTask(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	task, err := users[0].CreateTask("My first task", 1, 1, time.Now())
	if err != nil {
		t.Error(err, "Cannot create task")
	}
	if task.User.ID != users[0].ID {
		t.Error("User not linked with task")
	}
}