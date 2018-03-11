package rate

import (
	"testing"
	"time"
)

var ti = Tier{
	Fixed: Format{
		Time:  "MONTHLY",
		Value: 11,
	},
	Variable: Format{
		Time:  "MONTHLY",
		Value: 7,
	},
}

func TestOccurrence(t *testing.T) {
	actual := occurrence(ti, 2, MONTHLY, time.Now())
	expected := float64(8)
	if actual != expected {
		t.Errorf("Expected Occurrence be %f and the result was %f", expected, actual)
	}
	actual = occurrence(ti, 2, YEARLY, time.Now())
	expected = float64(228)
	if actual != expected {
		t.Errorf("Expected Occurrence be %f and the result was %f", expected, actual)
	}
}
