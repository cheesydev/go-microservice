package pi

// Calculate an aproximation of the pi value with 5 decimal digits of precision
// using the a formula based on Leibniz series.
func LeibnizPi() float64 {

	// aprox. number of terms needed to achieve precision of 5 digits
	return LeibnizPiTerms(800000)
}

// Calculate pi using Leibniz's formula with a given number of terms.
func LeibnizPiTerms(terms int) float64 {

	var pi float64
	for i := 1; i <= terms; i++ {
		signal := 1.
		if i%2 == 0 {
			signal = -1.
		}

		// 1, 3, 5, 7 ...
		div := float64(((i * 2) - 1))

		// +(4/1), -(4/3), +(4/5), -(4/7) ...
		var term = signal * (4 / div)

		pi = pi + term
	}

	return pi
}
