package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"todo/controller"
	"todo/log"
	"todo/middleware"
	"todo/model"
	"todo/router"
)

var dbCheck = flag.Bool("db_check", false, "Specify whether the application should check if it can connect to the database on startup.")
var action = flag.String("a", "", "Specify a utility action to execute, such as executing a SQL script, rather than running the main application.")
var dbName = flag.String("d", "todo", "Specify the database name to use for an action.")
var userName = flag.String("U", "todo", "Specify the user name to use for an action.")
var file = flag.String("f", "", "Specify a file path to use for an action.")

func handleFlags() {
	flag.Parse()

	// check if database exists
	if *dbCheck == true {
		checkDatabase()
	}

	// run SQL functions
	if *action == "sql_functions" {
		runScript(*dbName, *userName, *file)
	}
}

func checkDatabase() {
	err := model.CheckDatabase()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	} else {
		log.Info("Database connection was successful...")
	}
}

func runScript(_dbName string, _userName string, _file string) {
	err := model.RunScript(*dbName, *userName, *file)
	if err != nil {
		log.Fatal("Error running SQL functions:", err)
	}

	os.Exit(0)
}

func main() {
	handleFlags()

	// register all controller routes
	controller.RegisterRoutes()

	// handle static assets (css, js, etc.)
	router.HandleStatic()

	// wrap our router with middleware to log each request
	handler := middleware.LogRequest(router.Router())

	fmt.Print("Listening on port 8000...\n")
	http.ListenAndServe(":8000", handler)
}
