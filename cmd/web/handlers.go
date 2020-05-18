package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Change the signature of the home handler so it is defined as a method against *application.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) // use the notFound() helper
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // use the serverError() helper
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err) // use the serverError() helper
	}
}

// Add a showSnippet handler function
// Change the signature of the showSnippet handler so it is defined as a method against *application.
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to convert it to an integer using the strconv.Atoi() function. If it can't be converted to an integer, or the value is less than 1, we return a 404 page not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // use the notFound() helper
		return
	}
	// Use the fmt.Fprintf() function to interpolate the id value with our response and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a createSnippet handler function
// Change the signature of the createSnippet handler so it is defined as a method against *application.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not. Note that http.MethodPost is a constant equal to the string "POST".
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed) // use the clientError() helper
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
