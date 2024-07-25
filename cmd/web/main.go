package main

import (
	"log"
	"net/http"
)

func main() {
	// initialize a new server(new server mux)
	mux := http.NewServeMux()

	//register handler with router
	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/home", homeHandler)

	mux.HandleFunc("/home/create", createHandler)

	fileserver := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	log.Println("Startint the server on 3000")

	// start a new web server
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
