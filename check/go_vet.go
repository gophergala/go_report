package check

type GoVet struct {
	Dir string
}

func (g GoVet) Name() string {
	return "go_vet"
}

// Percentage returns the percentage of .go files that pass go vet
func (g GoVet) Percentage() (float64, map[string][]string, error) {
	return GoTool(g.Dir, []string{"go", "tool", "vet"})
}
