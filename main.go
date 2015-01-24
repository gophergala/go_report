package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gophergala/go_report/check"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/home.html")
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	repo := r.FormValue("url")
	if !strings.HasPrefix(repo, "https://") {
		repo = "https://" + repo
	}
	dir := strings.TrimSuffix(repo, ".git")
	split := strings.Split(dir, "/")
	dir = fmt.Sprintf("repos/%s-%s", split[len(split)-2], split[len(split)-1])
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", repo, dir)
		if err := cmd.Run(); err != nil {
			http.Error(w, fmt.Sprintf("Could not run git clone: %v", err), 500)
			return
		}
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Could not stat dir: %v", err), 500)
		return
	} else {
		cmd := exec.Command("git", "-C", dir, "pull")
		if err := cmd.Run(); err != nil {
			http.Error(w, fmt.Sprintf("Could not pull repo: %v", err), 500)
			return
		}
	}

	type score struct {
		Name       string  `json:"name"`
		Percentage float64 `json:"percentage"`
	}
	type checksResp struct {
		Checks []score `json:"checks"`
	}

	resp := checksResp{}
	checks := []check.Check{check.GoFmt{Dir: dir},
		//check.GoVet{Dir: dir},
		check.GoLint{Dir: dir},
		check.GoCyclo{Dir: dir},
	}

	for _, c := range checks {
		p, err := c.Percentage()
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not run check: %v", err), 500)
			return
		}
		ch := score{c.Name(), p}
		resp.Checks = append(resp.Checks, ch)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(b)
}

func main() {
	if err := os.Mkdir("repos", 0755); err != nil && !os.IsExist(err) {
		log.Fatal("could not create repos dir: ", err)
	}
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/checks", checkHandler)
	fmt.Println("Running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
