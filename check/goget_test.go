package check

import "testing"

func TestGetPackage(t *testing.T) {
	path := "github.com/nsf/gocode"
	err := GetPackage(path)
	if err != nil {
		t.Error(err)
	}
}
