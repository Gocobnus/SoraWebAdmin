package demo

import (
	"fmt"
	"net/http"
)

func Handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go1!</h1>")
}
