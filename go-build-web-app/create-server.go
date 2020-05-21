package createwebapp

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CreateServer(project string) {
	str := fmt.Sprint(`
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
	`)

	nf, err := os.Create(project + "/main.go")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
