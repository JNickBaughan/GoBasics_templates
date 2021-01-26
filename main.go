package main


import (
	"log"
	"os"
	"text/template"
	"net/http"
	"github.com/gorilla/mux"
)

var tpl *template.Template

func initTpl() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func handleMainRoute(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleStandardOut(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(os.Stdout, nil)
	if(err != nil){
		log.Fatalln(err)
	}
}

func main() {
	
	router := mux.NewRouter().StrictSlash(true)

	initTpl()

	router.HandleFunc("/", handleMainRoute)

	router.HandleFunc("/standard-out", handleStandardOut)

	log.Fatal(http.ListenAndServe(":3000", router))
}