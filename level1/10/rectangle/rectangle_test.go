package rectangle

import "testing"

func TestCalcArea(t *testing.T) {
	result := calcArea(3, 10)
	if result != 30 {
		t.Errorf("for 3, 10 expected 30")
	}
}
