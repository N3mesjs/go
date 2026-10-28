package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		fmt.Fprintf(w, "Title of the book: %s, page: %s", vars["title"], vars["page"])
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}