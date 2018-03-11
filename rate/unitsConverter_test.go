package rate

import (
	"testing"
)

func TestExtractPrefix(t *testing.T) {
	// SI symbol prefixes should be extracted
	for _, sym := range symbols {
		for pf, _ := range siPrefix {
			actual := extractPrefix(pf + sym)
			expected := pf
			if actual != expected {
				t.Errorf("Expected Extract prefix of %s to be %s and the result is %+v", pf+sym, expected, actual)
			}
		}
	}
	// Not found prefixes should return the full unit
	actual := extractPrefix("UNKNOWN")
	expected := "UNKNOWN"
	if actual != expected {
		t.Errorf("Expected Extract prefix to be %s and the result is %+v", expected, actual)
	}
	// nil unit returns empty string
	actual = extractPrefix("")
	expected = ""
	if actual != expected {
		t.Errorf("Expected Extract prefix to be %s and the result is %+v", expected, actual)
	}
}

func TestDistance(t *testing.T) {
	// SI symbol returns distance to base unit
	for sym, value := range siPrefix {
		actual := distance(sym, "", siPrefix)
		expected := value
		if actual != expected {
			t.Errorf("Expected Distance siPrefix from %s to %s --> to be %f and the result is %+v", sym, "_", expected, actual)
		}
	}

	// BINARY symbol returns distance to base unit
	for sym, value := range binaryPrefix {
		actual := distance(sym, "", binaryPrefix)
		expected := value
		if actual != expected {
			t.Errorf("Expected Distance binaryPrefix from %s to %s --> to be %f and the result is %+v", sym, "_", expected, actual)
		}
	}

	// allPrefix (default) symbol returns distance to base unit
	for sym, value := range allPrefix {
		actual := distance(sym, "", allPrefix)
		expected := value
		if actual != expected {
			t.Errorf("Expected Distance allPrefix from %s to %s --> to be %f and the result is %+v", sym, "_", expected, actual)
		}
	}

	// returns nil if origin or destination are not found
	for sym, _ := range allPrefix {
		actual := distance(sym, "UNKNOWN", allPrefix)
		expected := float64(0)
		if actual != expected {
			t.Errorf("Expected Distance allPrefix from %s to %s --> to be %f and the result is %+v", sym, "_", expected, actual)
		}
		actual = distance("UNKNOWN", sym, allPrefix)
		expected = float64(0)
		if actual != expected {
			t.Errorf("Expected Distance allPrefix from %s to %s --> to be %f and the result is %+v", sym, "_", expected, actual)
		}
	}

	// SI symbol returns distance between symbols
	origin := []string{"K", "M"}
	finish := []string{"K", "M"}
	for _, x := range origin {
		for _, y := range finish {
			actual := distance(x, y, siPrefix)
			expected := siPrefix[x] / siPrefix[y]
			if actual != expected {
				t.Errorf("Expected Distance siPrefix from %s to %s --> to be %f and the result is %+v", x, y, expected, actual)
			}
		}
	}

	// BINARY symbol returns distance between symbols
	origin = []string{"Ki", "Mi"}
	finish = []string{"Ki", "Mi"}
	for _, x := range origin {
		for _, y := range finish {
			actual := distance(x, y, binaryPrefix)
			expected := binaryPrefix[x] / binaryPrefix[y]
			if actual != expected {
				t.Errorf("Expected Distance binaryPrefix from %s to %s --> to be %f and the result is %+v", x, y, expected, actual)
			}
		}
	}

	// Default symbols returns distance between symbols
	origin = []string{"K", "M"}
	finish = []string{"Ki", "Mi"}
	for _, x := range origin {
		for _, y := range finish {
			actual := distance(x, y, allPrefix)
			expected := allPrefix[x] / allPrefix[y]
			if actual != expected {
				t.Errorf("Expected Distance allPrefix from %s to %s --> to be %f and the result is %+v", x, y, expected, actual)
			}
		}
	}
}

func TestToUnit(t *testing.T) {
	// SI symbol returns value in base unit
	actual := toUnit(7, "", "", allPrefix)
	expected := float64(7)
	if actual != expected {
		t.Errorf("Expected toUnit SI_SIMBOL to be %f and the result is %f", expected, actual)
	}
	actual = toUnit(7, "KB", "", allPrefix)
	expected = float64(7000)
	if actual != expected {
		t.Errorf("Expected toUnit SI_SIMBOL to be %f and the result is %f", expected, actual)
	}

	// BINARY symbol returns value in base unit
	actual = toUnit(7, "KiB", "", binaryPrefix)
	expected = float64(7168)
	if actual != expected {
		t.Errorf("Expected toUnit SI_SIMBOL to be %f and the result is %f", expected, actual)
	}

	// SI symbol returns value in destination unit
	actual = toUnit(7, "MB", "KB", allPrefix)
	expected = float64(7000)
	if actual != expected {
		t.Errorf("Expected toUnit SI_SIMBOL to be %f and the result is %f", expected, actual)
	}

	// BINARY symbol returns value in destination unit
	actual = toUnit(7, "PiB", "TiB", binaryPrefix)
	expected = float64(7168)
	if actual != expected {
		t.Errorf("Expected toUnit binaryPrefix to be %f and the result is %f", expected, actual)
	}

	// SI symbol returns value in destination unit
	actual = toUnit(7, "PB", "TiB", allPrefix)
	expected = float64(6366.462912410498)
	if actual != expected {
		t.Errorf("Expected toUnit allPrefix to be %f and the result is %f", expected, actual)
	}
}
