package main

import (
    "fmt"
    "net/http"
)

func main(){
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello World! You have requested: %s\n", r.URL.Path)
        for key, val := range r.URL.Query(){
            fmt.Fprintln(w, "Parameters Query")
            fmt.Fprintln(w, key, val)
            fmt.Fprintln(w, r.URL.Query())
        }
    })

    /*
    fs := http.FileServer(http.Dir("assets/"))

    http.Handle("/static/", http.StripPrefix("/static", fs))
    */

    http.ListenAndServe(":80", nil)
}