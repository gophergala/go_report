package check

type Check interface {
	Name() string
	// Percentage returns the passing percentage of the check,
	// as well as a map of filename to output
	Percentage() (float64, []FileSummary, error)
}
