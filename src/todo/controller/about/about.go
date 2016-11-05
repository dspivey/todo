package about

import (
	"net/http"
	"todo/router"
)

// Load routes for this controller
func Load() {
	router.Get("/about", index)
}

// Index displays the About page
func index(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("About"))
}
