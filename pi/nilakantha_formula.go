package pi

// Calculate an aproximation of the pi value with 4 decimal digits of precision
// using the a formula based on Nilankantha's series.
func NilakanthaPi() float64 {

	// aprox. num. of terms to provide precision of 5 digitis after the dot
	return NilakanthaPiTerms(60)
}

// Calculate an aproximation of the pi value using the Nilankantha's series
// with the provided number of terms n.
func NilakanthaPiTerms(n int) float64 {

	pi := 3.
	for i := 1; i <= n; i++ {
		signal := 1.
		if i%2 == 0 {
			signal = -1.
		}

		// 2*3*4, 4*5*6, 6*7*8, 8*9*10 ...
		divisor := (i * 2) * (i*2 + 1) * (i*2 + 2)

		// 4/(2*3*4), -4/(4*5*6), 4/(6*7*8), -4/(8*9*10) ...
		term := 4. / (signal * float64(divisor))

		pi = pi + term
	}

	return pi
}
