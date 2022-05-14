package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type User struct {
	Name string
	Age uint16
	Money int16
	Avg_grades float64
	Happiness float64
	Hobbies []string
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money: %d", u.Name, u.Age, u.Money)
} 

func (u *User) setNewName(name string) {
	u.Name = name
} 

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{
		Name: "Bob",
		Age: 25,
		Money: -50,
		Avg_grades: 4.2,
		Happiness: 0.8,
		Hobbies: []string{"football", "chess", "girls"},
	}
	// bob.setNewName("Alex")
	// fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Контакты")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {

	handleRequest()
}