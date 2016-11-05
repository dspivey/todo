package login

import (
	"net/http"
	"todo/log"
	"todo/model"
	"todo/router"
	"todo/view"
)

// Load routes for this controller
func Load() {
	router.Get("/login", login)
	router.Get("/signup", signup)
	router.Post("/signup", signupAccount)
}

// GET /login
// show the login page
func login(rw http.ResponseWriter, req *http.Request) {
	views := []string{"shared/_login.layout", "login/login"}
	template := view.ParseTemplates(views...)
	if template == nil {
		rw.Write([]byte("Template 'login.html' not found!"))
	} else {
		template.Execute(rw, nil)
	}
}

// GET /signup
// Show the signup page
func signup(rw http.ResponseWriter, req *http.Request) {
	views := []string{"shared/_login.layout", "signup/signup"}
	view.RenderHTML(rw, nil, views...)
}

// POST /signup
// Create the user account
func signupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Danger(err, "Cannot parse form")
	}

	user := model.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}

	if err := user.Create(); err != nil {
		log.Danger(err, "Cannot create user.")
	}

	http.Redirect(writer, request, "/login", 302)
}
