package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	case "/contact":
		fmt.Fprint(w, "<h1>Contact Us</h1>")
	default:
		fmt.Fprint(w, "<h1>The gophers couldn't find what you were looking for. :-/</h1>")
	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
