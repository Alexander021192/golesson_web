package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Alexander021192/web/internal/tmplaction"
)

type art struct {
	Id       string
	Title    string
	Anons    string
	FullText string
}

func handleRequest() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", tmplaction.Index).Methods("GET")
	rtr.HandleFunc("/create/", tmplaction.Create).Methods("GET")
	rtr.HandleFunc("/save_art", tmplaction.SaveArt).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", tmplaction.ShowPost).Methods("GET")


	http.Handle("/", rtr)

	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
