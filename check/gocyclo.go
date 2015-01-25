package check

type GoCyclo struct {
	Dir       string
	Filenames []string
}

func (g GoCyclo) Name() string {
	return "gocyclo"
}

// Percentage returns the percentage of .go files that pass gofmt
func (g GoCyclo) Percentage() (float64, []FileSummary, error) {
	return GoTool(g.Dir, g.Filenames, []string{"gocyclo", "-over", "10"})
}
