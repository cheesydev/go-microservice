package pi

import (
	"math"
	"testing"
)

func TestLeibnizFormula(t *testing.T) {
	pi := LeibnizPi()
	if pi < 3 || pi > 3.5 {
		t.Errorf("Bad aproximation of pi value: %f\n", pi)
	}
}

const threshold = 0.001

// comparing float numbers is problematic, instead check if they are close
// enough
func almostEquals(a, b float64) bool {
	return math.Abs(a-b) < threshold
}

func TestLeibnizFormulaTerms(t *testing.T) {

	var pi float64
	var expected float64

	pi = LeibnizPiTerms(1)
	expected = (4. / 1)
	if !almostEquals(pi, expected) {
		t.Errorf("Bad pi aproximation for 1 term. Expected: %f, got %f\n",
			expected, pi)
	}

	pi = LeibnizPiTerms(2)
	expected = (4. / 1) - (4. / 3)
	if !almostEquals(pi, expected) {
		t.Errorf("Bad pi aproximation for 2 terms. Expected: %f, got %f\n",
			expected, pi)
	}

	pi = LeibnizPiTerms(3)
	expected = (4. / 1) - (4. / 3) + (4. / 5)
	if !almostEquals(pi, expected) {
		t.Errorf("Bad pi aproximation for 3 terms. Expected: %f, got %f\n",
			expected, pi)
	}
}
