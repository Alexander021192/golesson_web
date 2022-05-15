package main

import (
	"net/http"

	"github.com/Alexander021192/web/internal/tmplaction"
)

type art struct {
	Id       string
	Title    string
	Anons    string
	FullText string
}



func handleRequest() {
	http.HandleFunc("/", tmplaction.Index)
	http.HandleFunc("/create/", tmplaction.Create)
	http.HandleFunc("/save_art", tmplaction.SaveArt)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
