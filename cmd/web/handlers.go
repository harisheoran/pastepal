package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) rootHandler(w http.ResponseWriter, req *http.Request) {
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
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = temp.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) homeHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, req)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	w.Write([]byte("Home"))
}

func (app *application) createHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", 405)
		http.Error(w, "Method not allowed", 405)
		return
	}

	w.Write([]byte("Create"))
}
