package tmplaction

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Alexander021192/web/internal/app/commands"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	data, err := commands.ReadCsv("test.csv")
	if err != nil {
		panic(err)
	}

	artList := commands.CreateArtList(data)

	t.ExecuteTemplate(w, "index", artList)
}
