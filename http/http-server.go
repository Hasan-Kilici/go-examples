package main

import (
	"fmt"
	"net/http"
	"html/template"
	"time"
	"path"
)

func handle(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    name := ""
     if r.URL.Path == "/" {
    name = "index.html"
    } else {
    name = path.Base(r.URL.Path)
    }

  data := struct{
    Author string
  }{
    Author: "Hasan Kılıcı",
	}
	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error", err)
	}
}

func main() {
	fmt.Println("http server açıldı!")
	// serves all the static files in the static directory
	http.Handle(
		"/static/",
		 http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	// runs the server on any ip and handles the index
	http.HandleFunc("/", handle)
	http.ListenAndServe(":0", nil)
}
