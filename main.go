package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"text/template"
)

type frenchBulldog struct {
	Name string
	NickName string
}

var tpl *template.Template

func initTpl() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func handleMainRoute(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", "templates")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleMyDogRoute(w http.ResponseWriter, r *http.Request) {
	data := []frenchBulldog{
		frenchBulldog{
			Name: "Oliver",
			NickName: "Stinks",
		},
		frenchBulldog{
			Name: "Gibson",
			NickName: "Monster",
		},
		frenchBulldog{
			Name: "Finnegan",
			NickName: "Nugget",
		},
		}
	err := tpl.ExecuteTemplate(w, "myDogs.gohtml", data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleStandardOut(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	initTpl()

	router.HandleFunc("/", handleMainRoute)

	router.HandleFunc("/my-dogs", handleMyDogRoute)

	router.HandleFunc("/standard-out", handleStandardOut)

	log.Fatal(http.ListenAndServe(":3000", router))
}
