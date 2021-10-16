package main

import (
	"fmt"
	"net/http"
)

func HelloWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h2>Hello from %s</h2>", r.URL.Path[1:])
}

func AnotherWebFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Another Message from here</h1>")
}

func main() {
	http.HandleFunc("/another", AnotherWebFunc)
	http.HandleFunc("/", HelloWeb)
	http.ListenAndServe(":8000", nil)
}
