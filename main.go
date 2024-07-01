package main

import (
	"fmt"
	"net/http"
)

// handler for serving the HTML page
func servePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// handler for serving the "Hello World" message
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	http.HandleFunc("/", servePage)
	http.HandleFunc("/hello", sayHello)

	fmt.Println("Server is running on http://localhost:80")
	http.ListenAndServe(":80", nil)
}
