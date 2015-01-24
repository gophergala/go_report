package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	//repo := r.FormValue("repo")
	//checks := []check.Check{check.GoFmt{Dir: repo}}
	//type resp struct {
	//	Checks []check.Check
	//}
	//for _, c := range checks {
	//	//
	//}
	// clone the repo
	// run the checks
	// return the json
}

func main() {
	if err := os.Mkdir("repos", 0755); err != nil && !os.IsExist(err) {
		log.Fatal("could not create repos dir: ", err)
	}
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/checks", checkHandler)
	fmt.Println("Running on 127.0.0.1:8080...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
