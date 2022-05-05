package mathematic

import "testing"

const error = "Value expected: %v, received: %v"

func TestAverage(t *testing.T) {
	valueExpected := 7.28
	value := Average(7.2, 9.9, 6.1, 5.9)

	if value != valueExpected {
		t.Errorf(error, value, valueExpected)
	}
}
