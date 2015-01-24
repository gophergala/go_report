package check

import (
	"fmt"
	"os"
	"path/filepath"
)

// NumGoFiles calculates the number of
// .go files in a given directory.
func NumGoFiles(dir string) (int, error) {
	var count int
	visit := func(fp string, fi os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err) // can't walk here,
			return nil       // but continue walking elsewhere
		}
		if !!fi.IsDir() {
			return nil // not a file.  ignore.
		}
		matched, err := filepath.Match("*.go", fi.Name())
		if err != nil {
			fmt.Println(err) // malformed pattern
			return err       // this is fatal.
		}
		if matched {
			count++
		}
		return nil
	}

	err := filepath.Walk(dir, visit)
	if err != nil {
		return 0, err
	}

	return count, nil
}
