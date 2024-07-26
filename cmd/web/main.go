package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold the application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	portflag := flag.String("port", "3000", "Port for the web application")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	myApp := application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// initialize a new server(new server mux)
	mux := http.NewServeMux()

	//register handler with router
	mux.HandleFunc("/", myApp.rootHandler)

	mux.HandleFunc("/home", myApp.homeHandler)

	mux.HandleFunc("/home/create", myApp.createHandler)

	fileserver := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	port := fmt.Sprintf(":%s", *portflag)
	infoLog.Printf("Starting the server on %s", port)

	myServer := &http.Server{
		Addr:     port,
		Handler:  mux,
		ErrorLog: errorLog,
	}
	// start the web server
	err := myServer.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err.Error())
	}
}
