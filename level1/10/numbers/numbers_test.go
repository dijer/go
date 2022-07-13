package numbers

import "testing"

func TestTranslateNumber(t *testing.T) {
	result := translateNumber(952)

	if result != "9 сотен, 5 десятков, 2 единиц" {
		t.Errorf("for 912 expected: 9 сотен, 5 десятков, 2 единиц")
	}
}
