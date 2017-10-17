package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Print the URI to screen
	fmt.Fprintf(w, "URI: /%s", r.URL.Path[1:])
}

func main() {
	// Port to listen on
	port := ":8080"

	color.Green("Listening on port" + port)
	color.White("ctrl + c to exit")

	// Respond to all URIs
	http.HandleFunc("/", handler)

	// Listen for requests
	http.ListenAndServe(port, nil)
}
