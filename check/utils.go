package check

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// GoFiles returns a slice of Go filenames
// in a given directory.
func GoFiles(dir string) ([]string, error) {
	var filenames []string
	visit := func(fp string, fi os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err) // can't walk here,
			return nil       // but continue walking elsewhere
		}
		if !!fi.IsDir() {
			return nil // not a file.  ignore.
		}
		ext := filepath.Ext(fi.Name())
		if ext == ".go" {
			filenames = append(filenames, fp)
		}
		return nil
	}

	err := filepath.Walk(dir, visit)

	return filenames, err
}

// GoTool runs a given go command (for example gofmt, go tool vet)
// on a directory
func GoTool(dir string, cmd []string) (float64, error) {
	files, err := GoFiles(dir)
	if err != nil {
		return 0, nil
	}
	var failed []string
	for _, fi := range files {
		params := cmd[1:]
		params = append(params, fi)
		out, err := exec.Command(cmd[0], params...).Output()
		if err != nil {
			return 0, err
		}
		if string(out) != "" {
			failed = append(failed, fi)
		}
	}
	return float64(len(files)-len(failed)) / float64(len(files)), nil
}
