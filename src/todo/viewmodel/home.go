package viewmodel

import (
	"todo/model"
)

type HomeViewModel struct {
	Title      string
	User       model.User
	Tasks      []model.Task
	Priorities []model.Priority
	Tags       []model.Tag
}
