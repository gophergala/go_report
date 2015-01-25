package check

type GoFmt struct {
	Dir string
}

func (g GoFmt) Name() string {
	return "gofmt"
}

// Percentage returns the percentage of .go files that pass gofmt
func (g GoFmt) Percentage() (float64, []FileSummary, error) {
	return GoTool(g.Dir, []string{"gofmt", "-s", "-l"})
}
