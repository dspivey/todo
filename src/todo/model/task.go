package model

import (
	"time"

	"github.com/lib/pq"
)

type Task struct {
	TaskId     int
	Value      string
	PriorityId int
	Priority   *Priority
	StatusId   int
	Status     *Status
	CreatedAt  time.Time
	DueAt      pq.NullTime
	CompleteAt pq.NullTime
	UserId     int
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
func (user *User) CreateTask(value string, priorityId int, statusId int, dueAt time.Time) (task Task, err error) {
	statement := `insert into tasks (value, user_id, priority_id, status_id, created_at, due_at)
    values($1, $2, $3, (select s.status_id from status s where s.value = 'Incomplete') , $4, $5)
    returning task_id, value, priority_id, status_id, created_at, due_at`

	stmt, err := Database.Prepare(statement)
	if err != nil {
		return task, err
	}
	defer stmt.Close()

	// insert the new task and populate the "Task" object
	err = stmt.QueryRow(
		value, user.UserId, priorityId, time.Now(), dueAt,
	).Scan(
		&task.TaskId, &task.Value, &task.PriorityId, &task.StatusId, &task.CreatedAt, &task.DueAt,
	)
	if err != nil {
		return task, err
	}

	// add user to the task
	task.UserId = user.UserId
	task.User = user

	return task, err
}

// Tasks get tasks for a user
func (user *User) Tasks() (tasks []Task, err error) {
	statement := `SELECT t.task_id,t.value as "task_value",t.created_at,t.due_at,t.complete_at,
	p.priority_id, p.value as "priority_value",
	s.status_id, s.value as "status_value"
    FROM todo.tasks t
    INNER JOIN todo.users u ON t.user_id = u.user_id
	INNER JOIN todo.priorities p ON p.priority_id = t.priority_id
	INNER JOIN todo.status s ON s.status_id = t.status_id
    WHERE u.user_id = $1`

	rows, err := Database.Query(statement, user.UserId)
	if err != nil {
		return tasks, err
	}

	var priorityValue string
	var statusValue string

	for rows.Next() {
		task := Task{}
		err := rows.Scan(
			&task.TaskId,
			&task.Value,
			&task.CreatedAt,
			&task.DueAt,
			&task.CompleteAt,
			&task.PriorityId,
			&priorityValue,
			&task.StatusId,
			&statusValue,
		)
		if err != nil {
			return tasks, err
		}

		// add priority info to task
		task.Priority = &Priority{
			task.PriorityId,
			priorityValue,
		}

		// add status info to task
		task.Status = &Status{
			task.StatusId,
			statusValue,
		}

		// add user info to task
		task.UserId = user.UserId
		task.User = user

		tasks = append(tasks, task)
	}
	rows.Close()

	return tasks, err
}

// Tags returns all tags attached to the task
func (task *Task) Tags() (tags []Tag, err error) {
	statement := `select t.tag_id, t.value 
	from task_tags tt
	inner join tags t on tt.task_id = $1 and tt.tag_id = t.tag_id`

	rows, err := Database.Query(statement, task.TaskId)
	if err != nil {
		return tags, err
	}

	for rows.Next() {
		tag := Tag{}
		err := rows.Scan(
			&tag.TagId,
			&tag.Value,
		)
		if err != nil {
			return tags, err
		}

		tags = append(tags, tag)
	}
	rows.Close()

	return tags, err
}
