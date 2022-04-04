package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// OS Knowledge: stdout, stdin, stderr ---> File descriptior

// fmt --> stdout
// log --> stdout, file, remote adress

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	router := http.NewServeMux()
	router.HandleFunc("/", home)
	router.HandleFunc("/snippet", showSnippet)
	router.HandleFunc("/snippet/create", createSnippet)

	// Use the http.ListenAndServe() function to start a new web server. We pass
	// two parameters: the TCP network address to listen on (in this case ":4000
	// and the servemux we just created. If http.ListenAndServe() returns an err
	// we use the log.Fatal() function to log the error message and exit.
	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":4000", router) //4000. porta gelen http isteklerini router karşılayaccak. İlgili handler'a route edecek.
	log.Fatal(err)
}

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Handler func
	w.Write([]byte("Hello from Snippetbox"))
}

type DilaraWriter struct{}

func (s DilaraWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	dilaraWriter := DilaraWriter{}
	_ = dilaraWriter

	fmt.Fprintf(w, "Gelen ID: %d\n", id)
}

// curl -X POST http://localhost:4000/snippet/create
// curl -i -X POST http://localhost:4000/snippet/createSnippet
// -i HEADER'ları da gösterir
func createSnippet(w http.ResponseWriter, r *http.Request) { // Handler func
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "method not allowed", 405)
		return
	}
	w.Write([]byte("Create Snippet"))
}
