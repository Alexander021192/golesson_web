package tmplaction

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Alexander021192/web/internal/app/commands"
	"github.com/gorilla/mux"
)

func ShowPost(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/show_post.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	vars := mux.Vars(r)
	// fmt.Fprintf(w, "ID: %v\n", vars["id"])

	data, err := commands.ReadCsv("test.csv")
	if err != nil {
		panic(err)
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print("failed get id post")
	}

	art := commands.GetArt(data, id)

	t.ExecuteTemplate(w, "show_post", art)

}
