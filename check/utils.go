package check

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
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
		if fi.IsDir() {
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
func GoTool(dir string, command []string) (float64, error) {
	files, err := GoFiles(dir)
	if err != nil {
		return 0, nil
	}
	var failed []string
	for _, fi := range files {
		params := command[1:]
		params = append(params, fi)

		cmd := exec.Command(command[0], params...)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return 0, err
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			return 0, err
		}

		err = cmd.Start()
		if err != nil {
			return 0, err
		}

		out, err := ioutil.ReadAll(stdout)
		if err != nil {
			return 0, err
		}

		errout, err := ioutil.ReadAll(stderr)
		if err != nil {
			return 0, err
		}

		err = cmd.Wait()
		if exitErr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0

			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				// some commands exit 1 when files fail to pass (for example go vet)
				if status.ExitStatus() != 1 {
					return 0, err
				}
			}
		}

		if string(out) != "" {
			failed = append(failed, fi)
		}

		// go vet logs to stderr
		if string(errout) != "" {
			failed = append(failed, fi)
		}
	}

	return float64(len(files)-len(failed)) / float64(len(files)), nil
}
