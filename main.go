package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/gophergala/go_report/check"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving home page")
	if r.URL.Path[1:] == "" {
		http.ServeFile(w, r, "templates/home.html")
	} else {
		http.NotFound(w, r)
	}
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving " + r.URL.Path[1:])
	http.ServeFile(w, r, r.URL.Path[1:])
}

func orgRepoNames(url string) (string, string) {
	dir := strings.TrimSuffix(url, ".git")
	split := strings.Split(dir, "/")
	org := split[len(split)-2]
	repoName := split[len(split)-1]

	return org, repoName
}

func dirName(url string) string {
	org, repoName := orgRepoNames(url)

	return fmt.Sprintf("repos/src/github.com/%s/%s", org, repoName)
}

func clone(url string) error {
	org, _ := orgRepoNames(url)
	if err := os.Mkdir(fmt.Sprintf("repos/src/github.com/%s", org), 0755); err != nil && !os.IsExist(err) {
		return fmt.Errorf("could not create dir: %v", err)
	}
	dir := dirName(url)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", "--depth", "1", "--single-branch", url, dir)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("could not run git clone: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("could not stat dir: %v", err)
	} else {
		cmd := exec.Command("git", "-C", dir, "pull")
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("could not pull repo: %v", err)
		}
	}

	return nil
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	repo := r.FormValue("repo")
	url := repo
	if !strings.HasPrefix(url, "https://github.com/") {
		url = "https://github.com/" + url
	}

	err := clone(url)
	if err != nil {
		log.Println("ERROR: could not clone repo: ", err)
		http.Error(w, fmt.Sprintf("Could not clone repo: %v", err), 500)
		return
	}

	type score struct {
		Name          string              `json:"name"`
		FileSummaries []check.FileSummary `json:"file_summaries"`
		Percentage    float64             `json:"percentage"`
	}
	type checksResp struct {
		Checks  []score `json:"checks"`
		Average float64 `json:"average"`
		Files   int     `json:"files"`
		Issues  int     `json:"issues"`
	}

	resp := checksResp{}
	dir := dirName(url)
	filenames, err := check.GoFiles(dir)
	if err != nil {
		log.Println("ERROR: could not get filenames: ", err)
		http.Error(w, fmt.Sprintf("Could not get filenames: %v", err), 500)
		return
	}
	checks := []check.Check{check.GoFmt{Dir: dir, Filenames: filenames},
		check.GoVet{Dir: dir, Filenames: filenames},
		check.GoLint{Dir: dir, Filenames: filenames},
		check.GoCyclo{Dir: dir, Filenames: filenames},
	}

	ch := make(chan score)
	for _, c := range checks {
		go func(c check.Check) {
			p, out, err := c.Percentage()
			if err != nil {
				log.Printf("ERROR: (%s) %v", c.Name(), err)
				//http.Error(w, fmt.Sprintf("Could not run check %v: %v\r\n%v", c.Name(), err, out), 500)
				//return
			}
			s := score{c.Name(), out, p}
			ch <- s
		}(c)
	}

	var avg float64
	var issues = make(map[string]bool)
	for i := 0; i < len(checks); i++ {
		s := <-ch
		resp.Checks = append(resp.Checks, s)
		avg += s.Percentage
		for _, fs := range s.FileSummaries {
			issues[fs.Filename] = true
		}
	}

	avg = avg / float64(len(checks))
	resp.Average = avg
	resp.Files = len(filenames)
	resp.Issues = len(issues)

	b, err := json.Marshal(resp)
	if err != nil {
		log.Println("ERROR: could not marshal json:", err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(b)
}

func reportHandler(w http.ResponseWriter, r *http.Request, org, repo string) {
	http.ServeFile(w, r, "templates/home.html")
}

func makeReportHandler(fn func(http.ResponseWriter, *http.Request, string, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validPath := regexp.MustCompile(`^/report/([a-zA-Z0-9\-_]+)/([a-zA-Z0-9\-_]+)$`)

		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[1], m[2])
	}
}

func main() {
	if err := os.MkdirAll("repos/src/github.com", 0755); err != nil && !os.IsExist(err) {
		log.Fatal("ERROR: could not create repos dir: ", err)
	}

	http.HandleFunc("/assets/", assetsHandler)
	http.HandleFunc("/checks", checkHandler)
	http.HandleFunc("/report/", makeReportHandler(reportHandler))
	http.HandleFunc("/", homeHandler)

	fmt.Println("Running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
