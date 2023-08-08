package main

import (
	"html/template"
	"log"
	"net/http"
)

// main page
func Home(w http.ResponseWriter, r *http.Request) {

	// add check to fully qualified URL
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("../../ui/html/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internet server error", 500)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internet server error", 500)

	}

}

// show resume
func ShowResume(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("there is a resume"))
}
