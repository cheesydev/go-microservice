package pi

// Calculate and aproximate value of Pi
type PiCalculator interface {

	// Calculate an aproximation of the pi value with at least 5 digits of
	// precision.
	CalculatePi() float64
}
