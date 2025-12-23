package main

import (
	"fmt"
	"net/http"
)

func main() {
	homePage := "./home.html"
	aboutPage := `
  <h1>About Page</h1>
  <p>This page contains content about this web service</p>
  `

	// Define a route and its handler to print some text
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Web services are easy with Go!")
	})

	// Serve a static home page
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, homePage)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, aboutPage)
	})

	// Start the web server on port 3000
	http.ListenAndServe(":3000", nil)
}
