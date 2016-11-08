package viewmodel

import (
	"todo/model"
)

type HomeViewModel struct {
	Title string
	User  model.User
	Tasks []model.Task
}
