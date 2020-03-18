package pi

func LeibnizPi() float64 {
	// it takes ~0.6s
	return LeibnizPiTerms(500000000)
}

// calculate pi using Leibniz's formula with a given number of terms
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
