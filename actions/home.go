package actions

import (
	"net/http"
	"fmt"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Shop!")
}
