package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a new command-line flag with the name 'addr', a default value of ":4000" and some short help text explaining what the flag controls. The value of the flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line flag. This reads in the command-line flag value and assigns it to the addr variable. You need to call this *before* you use the addr variable otherwise it will always contain the default value of ":4000". If any errors are encountered during parsing the application will be terminated.
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Use the http.ListenAndServe() function to start a new web server. We pass in two parameters: the TCP network address to listen on (in this case ":4000") and the servemux we just created. If http.ListenAndServe() returns an error we use the log.Fatal() function to log the error message and exit. Note that any error returned by http.ListenAndServe() is always non-nil.
	log.Println("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
