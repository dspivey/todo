package router

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()
}

// Router returns the router instance.
func Router() *mux.Router {
	return router
}

func getPublicDir() (path string) {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "src", "todo", "public")
}

// HandleStatic handles serving static content
func HandleStatic() {
	fileServer := http.FileServer(http.Dir(getPublicDir()))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
}

// Delete is a shortcut for router.Handle("DELETE", path, handle).
func Delete(path string, fn http.HandlerFunc) {
	router.HandleFunc(path, fn).Methods("DELETE")
}

// Get is a shortcut for router.Handle("GET", path, handle).
func Get(path string, fn http.HandlerFunc) {
	router.HandleFunc(path, fn).Methods("GET")
}

// Post is a shortcut for router.Handle("Post", path, handle).
func Post(path string, fn http.HandlerFunc) {
	router.HandleFunc(path, fn).Methods("POST")
}
