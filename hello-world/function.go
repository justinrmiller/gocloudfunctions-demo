package helloworld

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// F prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func F(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello World!")
		return
	}
	if d.Message == "" {
		fmt.Fprint(w, "Hello World!")
		return
	}
	fmt.Fprint(w, d.Message)
}
