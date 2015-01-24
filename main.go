package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func cloneHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Running on 127.0.0.1:8080...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
