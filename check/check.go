package check

type Check interface {
	Name() string
	// Percentage returns the passing percentage of the check
	Percentage() (float64, error)
}
