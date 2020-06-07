package index

import (
	"fmt"
	"net/http"
)

//Handler handles [/]
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
