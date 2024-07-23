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

	log.Println("Startint the server on 3000")

	// start a new web server
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

// regular go function
func rootHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	w.Write([]byte("Root"))
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Home"))
}

func createHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		w.Header().Add("Cache-Control", "public")
		w.Header().Add("Cache-Control", "max-age=31536000")
		w.Header().Del("Cache-Control")
		w.Header().Get("Cache-Control")
		http.Error(w, "Method not allowed", 405)
		w.Write([]byte("Method not allowed"))
		return
	}

	w.Write([]byte("Create"))
}
