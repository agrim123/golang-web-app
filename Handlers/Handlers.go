package Handlers

import (
	"../Models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Home")
}

func NameHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")
	fmt.Fprintln(rw, "Hello", name)
}

func BookJSONHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	book := Models.Book{"Building Web Apps with Go", "Jeremy Saenz"}

	js, err := json.Marshal(book)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func BookHTMLHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	book := Models.Book{"Building Web Apps with Go", "Jeremy Saenz"}

	fp := path.Join("templates", "books.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(rw, book); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func UsersHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	emails := []string{"email1", "email2"}
	user := Models.User{1, "User 1", emails}

	fmt.Fprintln(rw, user.String())
}
