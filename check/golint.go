package check

type GoLint struct {
	Dir string
}

func (g GoLint) Name() string {
	return "golint"
}

// Percentage returns the percentage of .go files that pass golint
func (g GoLint) Percentage() (float64, error) {
	return GoTool(g.Dir, []string{"golint"})
}
