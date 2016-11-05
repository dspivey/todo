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
	router.Get("/logout", logout)
	router.Get("/signup", signup)
	router.Post("/signup", signupAccount)
	router.Post("/authenticate", authenticate)
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

// POST /authenticate
// Authenticate the user given the email and password
func authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := model.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		log.Danger(err, "Cannot find user")
	}
	if user.Password == model.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/home", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}

// GET /logout
// Logs the user out
func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		log.Warning(err, "Failed to get cookie")
		session := model.Session{UUID: cookie.Value}
		session.DeleteByUUID()
	}

	http.Redirect(writer, request, "/", 302)
}
