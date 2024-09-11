package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func main() {
	fs := http.FileServer(http.Dir("resources"))
	http.Handle("/resources/", http.StripPrefix("/resources", fs))
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/login/", LoginPage)
	http.ListenAndServe("localhost:8080", nil)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("templates/LoginPage.html"))
	tmpl.Execute(w, nil)
}
func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("templates/MainPage.html"))
	tmpl.Execute(w, nil)
}
