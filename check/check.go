package check

type Check interface {
	// Percentage returns the passing percentage of the check
	Percentage() (float64, error)
}
