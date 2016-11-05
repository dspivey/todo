package controller

import (
	"todo/controller/about"
	"todo/controller/home"
	"todo/controller/login"
)

// RegisterRoutes loads the routes for each controller.
func RegisterRoutes() {
	about.Load()
	home.Load()
	login.Load()
}
