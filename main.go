package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"text/template"
	"strings"
)

type frenchBulldog struct {
	Name string
	NickName string
}

func getData() []frenchBulldog {
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
	return data
}

var templateFunctions = template.FuncMap{
	"ToUpper": strings.ToUpper,
}

var tpl *template.Template

func initTpl() {
	//tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	tpl = template.Must(template.New("").Funcs(templateFunctions).ParseGlob("templates/*.gohtml"))
}

func handleMainRoute(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", "templates")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleMyDogRoute(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "myDogs.gohtml", getData())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleMyDogUpperRoute(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "myDogsFunc.gohtml", getData())
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

	router.HandleFunc("/my-dogs-upper", handleMyDogUpperRoute)

	router.HandleFunc("/standard-out", handleStandardOut)

	log.Fatal(http.ListenAndServe(":3000", router))
}
