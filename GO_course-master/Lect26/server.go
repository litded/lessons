package main

import (
	"net/http"
	"text/template"
)

type Post struct {
	Title   string
	Content string
	Author  string
}

func ViewPost(w http.ResponseWriter, r *http.Request) {
	post := Post{"Go", "Go is super language!", "Bob"}
	tmpl := template.Must(template.New("new_template").Parse(`<div>
		<h1>{{ .Title}}</h1>
		<p>{{.Content}}</p>
		<p>{{ .Author}}</p>
	</div>`))
	tmpl.Execute(w, post)
}

func HelloWithTemplate(w http.ResponseWriter, r *http.Request) {
	message := "Hello from string"
	tmpl, _ := template.New("templator9000").Parse("<h1>Kek {{ .}}</h1>")
	tmpl.Execute(w, message)
}

func Templator(w http.ResponseWriter, r *http.Request) {
	post := Post{"Go", "Go is super language!", "Bob"}
	tmpl, _ := template.ParseFiles("templates/base.html")
	tmpl.Execute(w, post)
}

func main() {
	http.HandleFunc("/", HelloWithTemplate)
	http.HandleFunc("/post", ViewPost)
	http.HandleFunc("/temp", Templator)
	http.ListenAndServe(":8000", nil)
}
