package main

import (
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var tmpl *template.Template

func main() {
	fs := http.FileServer(http.Dir("resources"))
	http.Handle("/resources/", http.StripPrefix("/resources", fs))
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/login/", LoginPage)
	http.ListenAndServe("localhost:8070", nil)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("Templates/LoginPage.html"))
	tmpl.Execute(w, nil)
}
func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("Templates/MainPage.html"))
	tmpl.Execute(w, nil)
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pass, err := bcrypt.GenerateFromPassword([]byte(r.Form["Password"][0]), 10)
	if err != nil{
		panic(err)
	}
	var user = User{0, r.Form["Login"][0], pass}
}
