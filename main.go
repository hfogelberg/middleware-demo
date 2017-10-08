package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/about", aboutHandler)

	adm := r.PathPrefix("/admin/").Subrouter()
	adm.HandleFunc("/index", adminHandler)
	adm.HandleFunc("/create", createHandler)

	mux := http.NewServeMux()
	mux.Handle("/", r)
	mux.Handle("/admin/", negroni.New(
		negroni.HandlerFunc(adminMiddleware),
		negroni.Wrap(r),
	))

	n := negroni.Classic()
	n.UseHandler(mux)
	http.ListenAndServe(":3000", n)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Index</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Admin</h1>")
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Create</h1>")
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Edit</h1>")
}

func adminMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Middleware has checked that everything is OK!")
	next(w, r)
}
