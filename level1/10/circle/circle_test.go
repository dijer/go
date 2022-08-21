package circle

import "testing"

func TestCalcDiameter(t *testing.T) {
	result := calcDiameter(10)
	if result != 3.5682482323055424 {
		t.Errorf("for S=10 expected diameter 3.5682482323055424")
	}
}

func TestCalcCircumference(t *testing.T) {
	result := calcCircumference(10)
	if result != 31.41592653589793 {
		t.Errorf("for diameter 10 expected circumference 31.41592653589793")
	}
}
