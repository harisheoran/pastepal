package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// regular go function
func rootHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	temp, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = temp.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, req)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	w.Write([]byte("Home"))
}

func createHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", 405)
		http.Error(w, "Method not allowed", 405)
		return
	}

	w.Write([]byte("Create"))
}
