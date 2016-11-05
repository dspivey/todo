package main

import (
	"fmt"
	"net/http"
	"todo/controller"
	"todo/log"
	"todo/middleware"
	"todo/model"
	"todo/router"
)

func main() {
	// check if database exists
	err := model.CheckDatabase()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// register all controller routes
	controller.RegisterRoutes()

	// handle static assets (css, js, etc.)
	router.HandleStatic()

	// wrap our router with middleware to log each request
	handler := middleware.LogRequest(router.Router())

	fmt.Print("Listening on port 8000...\n")
	http.ListenAndServe(":8000", handler)
}
