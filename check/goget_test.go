package check

import "testing"

func TestGetPackage(t *testing.T) {
	//working example
	path := "github.com/nsf/gocode"
	err := GetPackage(path)
	if err != nil {
		t.Error(err)
	}

	path = "githubcom/nsf/gocode"
	err = GetPackage(path)
	if err == nil {
		t.Error("Should fail for non-existent repo")
	}

}
