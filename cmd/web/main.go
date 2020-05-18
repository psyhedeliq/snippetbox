package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Use log.New() to create a logger for writing information messages. This takes three parameters: the destination to write the logs to (os.Stdout), a string prefix for message (INFO followed by a tab), and flags to indicate what additional information to include (local date and time). Note that the flags are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages in the same way, but use stderr as the destination and use the log.Lshortfile flag to include the relevant file name and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// Write messages using the two new loggers, instead of the standard logger.
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
