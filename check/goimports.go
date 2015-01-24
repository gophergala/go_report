package check

type GoImports struct {
	Dir string
}

func (g GoImports) Name() string {
	return "goimports"
}

// Percentage returns the percentage of .go files that pass goimports
func (g GoImports) Percentage() (float64, map[string][]string, error) {
	return GoTool(g.Dir, []string{"goimports", "-l"})
}
