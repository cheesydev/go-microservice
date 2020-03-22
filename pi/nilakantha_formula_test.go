package pi

import (
	"testing"
)

func TestNilakanthaFormula(t *testing.T) {

	pi := 3.141592
	precision := 0.000001
	lower := pi - precision
	higher := pi + precision

	var p PiCalculator = Nilakantha{}
	value := p.CalculatePi()

	if value < lower || value > higher {
		t.Errorf("Bad aproximation of pi value. "+
			"Expected value between %f and %f, got: %f\n", lower, higher, value)
	}
}

func TestNilakanthaFormulaTerms(t *testing.T) {

	var pi float64
	var expected float64

	pi = NilakanthaPiTerms(1)
	expected = 3. + (4. / (2 * 3 * 4))
	if !almostEquals(pi, expected) {
		t.Errorf("Bad aproximation pi with 1 term. "+
			"Expected: %f, got: %f\n", expected, pi)
	}

	pi = NilakanthaPiTerms(2)
	expected = 3. + (4. / (2 * 3 * 4)) - (4. / (4 * 5 * 6))
	if !almostEquals(pi, expected) {
		t.Errorf("Bad aproximation pi with 2 term. "+
			"Expected: %f, got: %f\n", expected, pi)
	}

	pi = NilakanthaPiTerms(3)
	expected = 3. + (4. / (2 * 3 * 4)) - (4. / (4 * 5 * 6)) + (4. / (6 * 7 * 8))
	if !almostEquals(pi, expected) {
		t.Errorf("Bad aproximation pi with 3 term. "+
			"Expected: %f, got: %f\n", expected, pi)
	}
}
