package model

import (
	"time"

	"github.com/lib/pq"
)

type Task struct {
	TaskId     int
	Value      string
	PriorityId int
	Priority   string
	StatusId   int
	Status     string
	CreatedAt  time.Time
	DueAt      pq.NullTime
	CompleteAt pq.NullTime
	User       *User
}

const dateFormatString = "Jan 2 at 3:04pm"

// CreatedAtDate formats the CreatedAt field for better display
func (task *Task) CreatedAtDate() string {
	return task.CreatedAt.Format(dateFormatString)
}

// DueAtDate formats the CreatedAt field for better display
func (task *Task) DueAtDate() string {
	if task.DueAt.Valid {
		return task.DueAt.Time.Format(dateFormatString)
	}

	return ""
}

// CompleteAtDate formats the CreatedAt field for better display
func (task *Task) CompleteAtDate() string {
	if task.CompleteAt.Valid {
		return task.CompleteAt.Time.Format(dateFormatString)
	}

	return ""
}

// CreateTask creates a new task
func (user *User) CreateTask(value string, priority int, status int, dueAt time.Time) (task Task, err error) {
	statement := `insert into tasks (value, priority_id, status_id, created_at, due_at)
    values($1, $2, $3, $4, $5)
    returning task_id, value as "task_value", priority_id, status_id, created_at, due_at`

	stmt, err := Database.Prepare(statement)
	if err != nil {
		return task, err
	}
	defer stmt.Close()

	// insert the new task and populate the "Task" object
	err = stmt.QueryRow(value, priority, status, time.Now(), dueAt).Scan(&task.TaskId, &task.Value, &task.PriorityId, &task.StatusId, &task.CreatedAt, &task.DueAt)
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

func insertUserTask(user *User, task Task) (err error) {
	// add entry to user_tasks
	statement := `insert into user_tasks (task_id, user_id) values($1, $2)`
	stmt, err := Database.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	stmt.QueryRow(task.TaskId, user.UserId)

	return err
}

// Tasks get tasks for a user
func (user *User) Tasks() (tasks []Task, err error) {
	statement := `SELECT t.task_id, t.value, t.priority_id, t.status_id,
	t.created_at, t.due_at, t.complete_at, p.value, s.value
    FROM todo.user_tasks ut
	INNER JOIN todo.tasks t ON ut.task_id = t.task_id
    INNER JOIN todo.users u ON ut.user_id = u.user_id
	INNER JOIN todo.priorities p ON p.priority_id = t.priority_id
	INNER JOIN todo.status s ON s.status_id = t.status_id
    WHERE u.user_id = $1`

	rows, err := Database.Query(statement, user.UserId)
	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		task := Task{}
		if err = rows.Scan(&task.TaskId, &task.Value, &task.PriorityId, &task.StatusId, &task.CreatedAt, &task.DueAt, &task.CompleteAt, &task.Priority, &task.Status); err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	rows.Close()

	return tasks, err
}
