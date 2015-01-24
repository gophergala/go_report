package check

import (
	"fmt"
	"os"
	"path/filepath"
)

// GoFiles returns a slice of FileInfo
// for .go files in a given directory.
func GoFiles(dir string) ([]os.FileInfo, error) {
	var files []os.FileInfo
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
			files = append(files, fi)
		}
		return nil
	}

	err := filepath.Walk(dir, visit)

	return files, err
}
