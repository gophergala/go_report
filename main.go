package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gophergala/go_report/check"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	repo := r.FormValue("repo")
	if !strings.HasPrefix(repo, "https://") {
		repo = "https://" + repo
	}
	dir := strings.TrimSuffix(repo, ".git")
	split := strings.Split(dir, "/")
	dir = fmt.Sprintf("repos/%s", split[len(split)-1])
	cmd := exec.Command("git", "clone", repo, dir)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	err := cmd.Wait()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	checks := []check.Check{check.GoFmt{Dir: dir}}
	type check struct {
		Name       string  `json:"name"`
		Percentage float64 `json:"percentage"`
	}
	type checksResp struct {
		checks []check `json:"checks"`
	}

	resp := checksResp{}
	for _, c := range checks {
		p, err := c.Percentage()
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		ch := check{c.Name(), p}
		resp.checks = append(resp.checks, ch)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write(b)
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
