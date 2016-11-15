package viewmodel

import (
    "todo/model"
)

type TasksViewModel struct {
    User    model.User
    Tasks   []model.Task
}