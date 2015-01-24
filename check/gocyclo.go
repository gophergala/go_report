package check

type GoCyclo struct {
	Dir string
}

func (g GoCyclo) Name() string {
	return "gocyclo"
}

// Percentage returns the percentage of .go files that pass gofmt
func (g GoCyclo) Percentage() (float64, map[string][]string, error) {
	return GoTool(g.Dir, []string{"gocyclo", "-over", "10"})
}
