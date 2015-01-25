package check

type GoVet struct {
	Dir       string
	Filenames []string
}

func (g GoVet) Name() string {
	return "go_vet"
}

// Percentage returns the percentage of .go files that pass go vet
func (g GoVet) Percentage() (float64, []FileSummary, error) {
	return GoTool(g.Dir, g.Filenames, []string{"go", "tool", "vet"})
}
