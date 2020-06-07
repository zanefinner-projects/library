package index

import (
	"fmt"
	"net/http"
)

//Handler handles [/]
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Routes:
	[/] ⟹ show all routes
	POST: [/book] ⟹ create book
	GET [/books] ⟹ show books
	DELETE: [/book/:id] ⟹ delete book
	PUT: [/book/id] ⟹ edit book`)
}
