package main

import (
"html/template"
"net/http"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
uploadTemplate := template.Must(template.ParseFiles("welcome.gtpl"))
uploadTemplate.Execute(w, nil)
}



func main() {
http.HandleFunc("/", indexHandle)
http.ListenAndServe(":9090", nil)
}