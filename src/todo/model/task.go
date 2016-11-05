package model

import (
	"time"
)

type Task struct {
	ID         	int
	Value      	string
	PriorityID 	int
	Priority   	string
	StatusID   	int
	Status     	string
	CreatedAt  	time.Time
	DueAt      	time.Time
	CompleteAt 	time.Time
	User		*User
}

const dateFormatString = "Jan 2, 2006 at 3:04pm"

// CreatedAtDate formats the CreatedAt field for better display
func (task *Task) CreatedAtDate() string {
	return task.CreatedAt.Format(dateFormatString)
}

// DueAtDate formats the CreatedAt field for better display
func (task *Task) DueAtDate() string {
	return task.DueAt.Format(dateFormatString)
}

// CompleteAtDate formats the CreatedAt field for better display
func (task *Task) CompleteAtDate() string {
	return task.CompleteAt.Format(dateFormatString)
}

// CreateTask creates a new task
func (user *User) CreateTask(value string, priority int, status int, dueAt time.Time)(task Task, err error) {
    statement := `insert into tasks (value, priority_id, status_id, created_at, due_at)
    values($1, $2, $3, $4, $5)
    returning id, value, priority_id, status_id, created_at, due_at`

    stmt, err := Database.Prepare(statement)
    if err != nil {
        return task, err
    }
    defer stmt.Close()

    // insert the new task and populate the "Task" object
	err = stmt.QueryRow(value, priority, status, time.Now(), dueAt).Scan(&task.ID, &task.Value, &task.PriorityID, &task.StatusID, &task.CreatedAt, &task.DueAt)
	if err != nil {
		return task, err
	}

	// insert record matching the current user to the new task
	err = insertUserTask(user, task)
	if err != nil {
		return task, err
	}

	// add user to the task
	task.User = user

    return task, err
}

func insertUserTask(user *User, task Task)(err error){
	// add entry to user_tasks
	statement := `insert into user_tasks (task_id, user_id) values($1, $2)`
	stmt, err := Database.Prepare(statement)
	if err != nil {
        return err
    }
    defer stmt.Close()

	stmt.QueryRow(task.ID, user.ID)
	
	return err 
}

// Tasks get tasks for a user
func (user *User) Tasks() (tasks []Task, err error) {
    statement := `SELECT id,value,priority_id,status_id,created_at,due_at
    FROM user_tasks ut
    INNER JOIN users u ON ut.user_id = u.id
    WHERE u.id = $1`

	rows, err := Database.Query(statement, user.ID)
	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		task := Task{}
		if err = rows.Scan(&task.ID, &task.Value, &task.PriorityID, &task.StatusID, &task.CreatedAt, &task.DueAt, &task.CompleteAt); err != nil {
            return tasks, err
		}
		tasks = append(tasks, task)
	}
	rows.Close()
    
	return tasks, err
}