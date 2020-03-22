package pi

import (
	"math"
	"testing"
)

func TestLeibnizCalculater(t *testing.T) {
	TestLeibnizFormula(t)
}

func TestLeibnizFormula(t *testing.T) {

	pi := 3.141592
	precision := 0.000001
	lower := pi - precision
	higher := pi + precision

	var p PiCalculator = Leibniz{}
	value := p.CalculatePi()

	if value < lower || value > higher {
		t.Errorf("Bad aproximation of pi value. "+
			"Expecting value between %f and %f, got: %f\n",
			lower,
			higher,
			value)
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
