package actions

import (
	"fmt"
	"net/http"

	"github.com/AndriiOmelianenko/shop-api/types"
)

// GetVersion default implementation.
// curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/version
func GetVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Shop API version:", types.Version)

}
