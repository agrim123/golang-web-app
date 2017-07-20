package Handlers

import (
	"../Database"
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
	book := Models.Book{"Building Web Apps with Go", "Some Author"}

	js, err := json.Marshal(book)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func BookHTMLHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	book := Models.Book{"Building Web Apps with Go", "Some Author"}

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
	user := Models.User{Name: "Agrim"}
	Database.Insert(user)
}
