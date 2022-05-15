package tmplaction

import (
	"fmt"
	"html/template"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/header.html", "templates/footer.html", "templates/create.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}
