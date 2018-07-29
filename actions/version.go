package actions

import (
	"fmt"
	"net/http"

	"github.com/AndriiOmelianenko/shop-api/types"
)

// GetVersion default implementation.
func GetVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Shop API version:", types.Version)

}
