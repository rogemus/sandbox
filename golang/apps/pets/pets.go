package main

import (
	"html/template"
	"log"
	"net/http"
)

type Pet struct {
	Name   string
	Sex    string
	Intact bool
	Age    string
	Breed  string
}

var dogs = []Pet{
	{
		Name:   "Jujube",
		Sex:    "Female",
		Intact: false,
		Age:    "10 months",
		Breed:  "German Shepherd/Pitbull",
	},
	{
		Name:   "Zephyr",
		Sex:    "Male",
		Intact: true,
		Age:    "13 years, 3 months",
		Breed:  "German Shepherd/Border Collie",
	},
}

func homePage(w http.ResponseWriter, r *http.Request) {
  tmpl, err := template.New("").ParseFiles("./ui/home.html", "./ui/base.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(w, "base", dogs)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8181", nil))
}
