package check

import (
	"encoding/json"
	"log"
	"os/exec"
)

func isOk(err error) bool {
	if err == nil {
		return true
	} else {
		return false
	}
}

type Package struct {
	Dir        string
	DepsErrors []*PackageError
}

// A PackageError describes an error loading information about a package.
type PackageError struct {
	ImportStack   []string // shortest path from package named on command line to this one
	Pos           string   // position of error
	Err           string   // the error itself
	isImportCycle bool     // the error is an import cycle
	hard          bool     // whether the error is soft or hard; soft errors are ignored in some places
}

func (p *Package) GetDependencies() error {
	depList := p.ListPendingDependencies()
	for _, dep := range depList {
		err := GetPackage(dep)
		if !isOk(err) {
			return err
		}
	}
	return nil
}

func (p *Package) ListPendingDependencies() []string {
	out, err := exec.Command("go", "list", "-json", p.Dir).Output()
	if !isOk(err) {
		log.Fatal(err)
	}
	pInfo := Package{}
	err = json.Unmarshal(out, &pInfo)
	if !isOk(err) {
		log.Fatal(err)
	}
	pList := []string{}
	for _, depErr := range pInfo.DepsErrors {
		pList = append(pList, depErr.ImportStack[len(depErr.ImportStack)-1])
	}
	return pList
}

func GetPackage(path string) error {
	_, err := exec.Command("go", "get", path).CombinedOutput()
	return err
}
