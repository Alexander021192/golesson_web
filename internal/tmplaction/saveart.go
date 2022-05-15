package tmplaction

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/Alexander021192/web/internal/app/commands"
)

func SaveArt(w http.ResponseWriter, r *http.Request) {

	data, _ := commands.ReadCsv("test.csv")

	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	fmt.Println(title, anons, full_text)

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprint(w, "не все данные ")
	} else {
		data, _ = commands.AddRow(data, []string{strconv.Itoa(len(data)), title, anons, full_text})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
