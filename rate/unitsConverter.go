package rate

import (
	"regexp"
)

func extractPrefix(unit string) string {
	var prefix []byte
	for _, sym := range symbols {
		re := regexp.MustCompile(`(.*)` + sym + `$`)
		if re.MatchString(unit) {
			prefix = re.ReplaceAll([]byte(unit), []byte("$1"))
			break
		}
	}
	if len(prefix) > 0 {
		return string(prefix)
	}
	return unit
}

func getPrefix(prefix string, otherPrefix string, list map[string]float64) float64 {
	orig, okOrig := list[prefix]
	if len(prefix) == 0 {
		orig = 1
	}
	dest, okDest := list[otherPrefix]
	if len(otherPrefix) == 0 {
		dest = 1
	}
	if (okDest || len(otherPrefix) == 0) && (okOrig || len(prefix) == 0) {
		return float64(orig / dest)
	}
	return float64(0)
}

func distance(prefix string, otherPrefix string, list map[string]float64) float64 {
	// Returns the distance and whether you need to divide or multiply
	// Check that the list of conversions exists or use the International Sistem SI
	if prefix == otherPrefix {
		return float64(1)
	}
	return getPrefix(prefix, otherPrefix, list)
}

func toUnit(value float64, unit string, destUnit string, prefixType map[string]float64) float64 {
	prefix := extractPrefix(unit)
	destPrefix := extractPrefix(destUnit)
	prefixDistance := distance(prefix, destPrefix, prefixType)
	if prefixDistance == float64(0) {
		return float64(0)
	}
	return (value * prefixDistance)
}
