package check

type GoFmt struct {
	Dir       string
	Filenames []string
}

func (g GoFmt) Name() string {
	return "gofmt"
}

// Percentage returns the percentage of .go files that pass gofmt
func (g GoFmt) Percentage() (float64, []FileSummary, error) {
	return GoTool(g.Dir, g.Filenames, []string{"gofmt", "-s", "-l"})
}
