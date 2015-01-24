package check

import (
	"fmt"
	"testing"
)

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

func TestListPendingDependencies(t *testing.T) {
	path := "./test/missing_dependencies"
	p := Package{Dir: path}

	dList := p.ListPendingDependencies()
	if len(dList) == 0 {
		t.Error("Dependency list should be non-empty")
	}
	for _, dep := range dList {
		fmt.Println(string(dep))
	}
}

func TestGetDependencies(t *testing.T) {
	path := "./test/missing_dependencies"
	p := Package{Dir: path}
	err := p.GetDependencies()
	if err != nil {
		t.Error(err)
	}

	dList := p.ListPendingDependencies()
	if len(dList) > 0 {
		t.Error("There should be no more dependencies")
	}
}
