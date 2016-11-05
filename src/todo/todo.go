package main

import (
	"fmt"
	"net/http"
	"todo/controller"
	"todo/middleware"
	"todo/router"
)

func main() {
	controller.RegisterRoutes()

	// wrap our router with middleware to log each request
	handler := middleware.LogRequest(router.Router())

	// handle static assets (css, js, etc.)
	router.HandleStatic()

	fmt.Print("Listening on port 8000...\n")
	http.ListenAndServe(":8000", handler)
}
