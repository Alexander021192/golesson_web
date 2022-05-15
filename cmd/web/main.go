package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"github.com/Alexander021192/web/internal/app/commands"
)

type art struct {
	Id       string
	Title    string
	Anons    string
	FullText string
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/header.html", "templates/footer.html", "templates/create.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

// func readCsv(filename string) ([][]string, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	csvReader := csv.NewReader(file)
// 	// csvReader.Comma = ';'

// 	data, err := csvReader.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// convert records to array of structs
// 	// ffList := createFfList(data)
// 	// return ffList, err
// 	return data, err
// }

// func addRow(data [][]string, row []string) ([][]string, error) {
// 	fmt.Println(row)
// 	data = append(data, row)
// 	fmt.Println(data)

// 	f, err := os.Create("test.csv")
// 	defer f.Close()

// 	if err != nil {
// 		log.Fatalln("failed to open file", err)
// 	}
// 	w := csv.NewWriter(f)
// 	err = w.WriteAll(data)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return data, err
// }

func save_art(w http.ResponseWriter, r *http.Request) {

	data, _ := comands.readCsv("test.csv")

	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprint(w, "не все данные ")
	} else {
		data, _ = comands.addRow(data, []string{strconv.Itoa(len(data)), title, anons, full_text})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func handleRequest() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create/", create)
	http.HandleFunc("/save_art", save_art)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}