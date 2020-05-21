
	package main

	import (
		"net/http"
		"text/template"
	)
	
	var tpl *template.Template
	
	func init() {
		tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	}
	
	func home(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, "index.gohtml")
	}
	
	func main() {
		fs := http.FileServer(http.Dir("assets/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))

		http.HandleFunc("/", home)
		http.ListenAndServe(":8080", nil)
	
	}
	