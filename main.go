package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>")

}

func cloneHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Running on 127.0.0.1:8080...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
