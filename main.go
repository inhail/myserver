package main

import (
"html/template"
"net/http"
"encoding/json"
"log"
)

type GPSdata struct {
	Latitude string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

var upload []GPSdata

func indexHandle(w http.ResponseWriter, r *http.Request) {
	uploadTemplate := template.Must(template.ParseFiles("welcome.gtpl"))
	uploadTemplate.Execute(w, nil)
}

func CreateGPSEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandle)
	router.HandleFunc("/json", CreateGPSEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090",router))
}