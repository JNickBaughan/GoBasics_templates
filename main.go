package main


import (
	"log"
	"os"
	"text/template"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func handleMainRoute(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if(err != nil){
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if(err != nil){
		log.Fatalln(err)
	}
	
	if err != nil {
		fmt.Fprintf(w, "Uh Oh... looks like something went wrong")
	}
	
	tpl.Execute(w, nil)
}

func main() {
	
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handleMainRoute)

	log.Fatal(http.ListenAndServe(":3000", router))
}