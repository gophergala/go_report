package check

import (
	"log"
	"os/exec"
)

//GetPackages go gets the packages for
// a given package
type GoPackage struct {
	Dir string
}

func (g *GoPackage) GetDependencies() error {
	return nil
}

func GetPackage(path string) error {
	out, err := exec.Command("go", "get", path).Output()
	log.Println(out)
	return err
}
