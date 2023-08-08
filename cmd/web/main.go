package main

import (
	"log"
	"net/http"
)

func main() {
	// initializing
	mux := http.NewServeMux()
	// home page
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/resume", ShowResume)

	// // conect static assets
	// fs := http.FileServer(http.Dir("tmp/"))

	// http.Handle("/tmp/", http.StripPrefix("/static/", fs))

	// create Port web-server
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
