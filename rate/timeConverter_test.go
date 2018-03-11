package rate

import (
	"testing"
	"time"
)

func TestCalculateDaysOfMonth(t *testing.T) {
	actual := float64(daysMonth(2018, time.February))
	expected := time.Date(2018, 3, 0, 0, 0, 0, 0, time.UTC).
		Sub(time.Date(2018, 2, 0, 0, 0, 0, 0, time.UTC)).Hours() / 24
	if actual != expected {
		t.Errorf("Expected Duration February 2018 be %+v days and the result is %+v", expected, actual)
	}

	actual = float64(daysMonth(2018, time.March))
	expected = time.Date(2018, 4, 0, 0, 0, 0, 0, time.UTC).
		Sub(time.Date(2018, 3, 0, 0, 0, 0, 0, time.UTC)).Hours() / 24
	if actual != expected {
		t.Errorf("Expected Duration March 2018 be %+v days and the result is %+v", expected, actual)
	}

	// Leap Year 2016
	actual = float64(daysMonth(2016, time.February))
	expected = time.Date(2016, 3, 0, 0, 0, 0, 0, time.UTC).
		Sub(time.Date(2016, 2, 0, 0, 0, 0, 0, time.UTC)).Hours() / 24
	if actual != expected && expected == 29 {
		t.Errorf("Expected Duration February 2018 be %+v days and the result is %+v", expected, actual)
	}
}

func TestCalculateDaysOfYear(t *testing.T) {

	// Leap Year 2016
	actual := float64(daysYear(2016))
	expected := time.Date(2016, time.December, 31, 0, 0, 0, 0, time.UTC).
		Sub(time.Date(2016, time.January, 0, 0, 0, 0, 0, time.UTC)).Hours() / 24
	if actual != expected && expected == 366 {
		t.Errorf("Expected Duration 2016 be %+v days and the result is %+v", expected, actual)
	}

	// Normal Year 2018
	actual = float64(daysYear(2018))
	expected = time.Date(2018, time.December, 31, 0, 0, 0, 0, time.UTC).
		Sub(time.Date(2018, time.January, 0, 0, 0, 0, 0, time.UTC)).Hours() / 24
	if actual != expected {
		t.Errorf("Expected Duration 2018 be %+v days and the result is %+v", expected, actual)
	}
}

func TestUnits(t *testing.T) {
	actual := MINUTELY.Seconds()
	expected := float64(60)
	if actual != expected {
		t.Errorf("Expected MINUTELY be %+v and the result is %+v", expected, actual)
	}

	actual = HOURLY.Seconds()
	expected = float64(3600)
	if actual != expected {
		t.Errorf("Expected HOURLY be %+v and the result is %+v", expected, actual)
	}

	actual = DAILY.Seconds()
	expected = float64(24 * 3600)
	if actual != expected {
		t.Errorf("Expected DAILY be %+v and the result is %+v", expected, actual)
	}

	actual = WEEKLY.Seconds()
	expected = float64(7 * 24 * 3600)
	if actual != expected {
		t.Errorf("Expected WEEKLY be %+v and the result is %+v", expected, actual)
	}

	actual = MONTHLY.Seconds()
	expected = float64(daysMonth(time.Now().Year(), time.Now().Month()) * 24 * 3600)
	if actual != expected {
		t.Errorf("Expected MONTHLY be %+v and the result is %+v", expected, actual)
	}

	actual = YEARLY.Seconds()
	expected = float64(daysYear(time.Now().Year()) * 24 * 3600)
	if actual != expected {
		t.Errorf("Expected YEARLY be %+v and the result is %+v", expected, actual)
	}
}

// number of intervals on this month

var TIME_VALUES = []time.Duration{
	0,
	15 * MINUTELY,
	45 * MINUTELY,
	1 * HOURLY,
	90 * MINUTELY,
	5 * HOURLY,
	1 * DAILY,
	36 * HOURLY,
	1 * WEEKLY,
	14112 * MINUTELY,
	30*MONTHLY - 1*time.Second,
}

func TestMinutely(t *testing.T) {
	interval := "MINUTELY"
	results := []float64{1, 15, 45, 60, 90, 300, 24 * 60, 15 * 24 * 6, 7 * 24 * 60, 14 * 7 * 24 * 6, 24 * 60 * MONTHLY.Seconds()}
	if len(results) != len(TIME_VALUES) {
		t.Errorf("Expected results amd TIME_VALUES same length")
	}

	start_t := time.Date(time.Now().Year(), time.Now().Month(), 0, 0, 0, 0, 0, time.UTC)
	var end_t time.Time
	var conversion float64
	for i, time_value := range TIME_VALUES {
		end_t = start_t.Add(time_value)
		if start_t.Month() != end_t.Month() {
			continue
		} else {
			conversion = numberOfIntervals(end_t.Sub(start_t), interval, time.Now(), 0, 0)
		}
		if conversion != results[i] {
			t.Errorf("Expected with %s for %f s to match %f, start: %s, end: %s, got %f",
				interval, time_value.Seconds(), results[i], start_t, end_t, conversion)
		}
	}
}

func TestHourly(t *testing.T) {
	interval := "HOURLY"
	results := []float64{1, 1, 1, 1, 2, 5, 24, 36, 168, 236, 24 * MONTHLY.Hours()}
	if len(results) != len(TIME_VALUES) {
		t.Errorf("Expected results amd TIME_VALUES same length")
	}

	start_t := time.Date(time.Now().Year(), time.Now().Month(), 0, 0, 0, 0, 0, time.UTC)
	var end_t time.Time
	var conversion float64
	for i, time_value := range TIME_VALUES {
		end_t = start_t.Add(time_value)
		if start_t.Month() != end_t.Month() {
			continue
		} else {
			conversion = numberOfIntervals(end_t.Sub(start_t), interval, time.Now(), 0, 0)
		}
		if conversion != results[i] {
			t.Errorf("Expected with %s for %f s to match %f, start: %s, end: %s, got %f",
				interval, time_value.Seconds(), results[i], start_t, end_t, conversion)
		}
	}
}

func TestDaily(t *testing.T) {
	interval := "DAILY"
	results := []float64{1, 1, 1, 1, 1, 1, 1, 2, 7, 10, float64(daysMonth(time.Now().Year(), time.Now().Month()))}
	if len(results) != len(TIME_VALUES) {
		t.Errorf("Expected results amd TIME_VALUES same length")
	}

	start_t := time.Date(time.Now().Year(), time.Now().Month(), 0, 0, 0, 0, 0, time.UTC)
	var end_t time.Time
	var conversion float64
	for i, time_value := range TIME_VALUES {
		end_t = start_t.Add(time_value)
		if start_t.Month() != end_t.Month() {
			continue
		} else {
			conversion = numberOfIntervals(end_t.Sub(start_t), interval, time.Now(), 0, 0)
		}
		if conversion != results[i] {
			t.Errorf("Expected with %s for %f s to match %f, start: %s, end: %s, got %f",
				interval, time_value.Seconds(), results[i], start_t, end_t, conversion)
		}
	}
}

func TestWeekly(t *testing.T) {
	interval := "WEEKLY"
	results := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 5}
	if len(results) != len(TIME_VALUES) {
		t.Errorf("Expected results amd TIME_VALUES same length")
	}

	start_t := time.Date(time.Now().Year(), time.Now().Month(), 0, 0, 0, 0, 0, time.UTC)
	var end_t time.Time
	var conversion float64
	for i, time_value := range TIME_VALUES {
		end_t = start_t.Add(time_value)
		if start_t.Month() != end_t.Month() {
			continue
		} else {
			conversion = numberOfIntervals(end_t.Sub(start_t), interval, time.Now(), 0, 0)
		}
		if conversion != results[i] {
			t.Errorf("Expected with %s for %f s to match %f, start: %s, end: %s, got %f",
				interval, time_value.Seconds(), results[i], start_t, end_t, conversion)
		}
	}
}

func TestMonthly(t *testing.T) {
	interval := "MONTHLY"
	results := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	if len(results) != len(TIME_VALUES) {
		t.Errorf("Expected results amd TIME_VALUES same length")
	}

	start_t := time.Date(time.Now().Year(), time.Now().Month(), 0, 0, 0, 0, 0, time.UTC)
	var end_t time.Time
	var conversion float64
	for i, time_value := range TIME_VALUES {
		end_t = start_t.Add(time_value)
		if start_t.Month() != end_t.Month() {
			continue
		} else {
			conversion = numberOfIntervals(end_t.Sub(start_t), interval, time.Now(), 0, 0)
		}
		if conversion != results[i] {
			t.Errorf("Expected with %s for %f s to match %f, start: %s, end: %s, got %f",
				interval, time_value.Seconds(), results[i], start_t, end_t, conversion)
		}
	}
}

func TestYearly(t *testing.T) {
	interval := "YEARLY"
	results := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	if len(results) != len(TIME_VALUES) {
		t.Errorf("Expected results amd TIME_VALUES same length")
	}

	start_t := time.Date(time.Now().Year(), time.Now().Month(), 0, 0, 0, 0, 0, time.UTC)
	var end_t time.Time
	var conversion float64
	for i, time_value := range TIME_VALUES {
		end_t = start_t.Add(time_value)
		if start_t.Month() != end_t.Month() {
			continue
		} else {
			conversion = numberOfIntervals(end_t.Sub(start_t), interval, time.Now(), 0, 0)
		}
		if conversion != results[i] {
			t.Errorf("Expected with %s for %f s to match %f, start: %s, end: %s, got %f",
				interval, time_value.Seconds(), results[i], start_t, end_t, conversion)
		}
	}
}

// calculating for a different month than current

var DIF_TIME_VALUES = []time.Duration{
	0,
	15 * MINUTELY,
	45 * MINUTELY,
	1 * HOURLY,
	90 * MINUTELY,
	5 * HOURLY,
	1 * DAILY,
	36 * HOURLY,
	1 * WEEKLY,
	14112 * MINUTELY,
	28*DAILY - 1*time.Second,
}

func TestHourlyDifferentMonth(t *testing.T) {
	interval := "HOURLY"
	results := []float64{1, 1, 1, 1, 2, 5, 24, 36, 168, 236, 28 * 24}
	if len(results) != len(DIF_TIME_VALUES) {
		t.Errorf("Expected results amd TIME_VALUES same length")
	}
	start_t := time.Date(time.Now().Year(), time.Now().Month(), 0, 0, 0, 0, 0, time.UTC)
	var end_t time.Time
	var conversion float64
	for i, time_value := range DIF_TIME_VALUES {
		end_t = start_t.Add(time_value)
		if start_t.Month() != end_t.Month() {
			continue
		} else {
			conversion = numberOfIntervals(end_t.Sub(start_t), interval, time.Date(2017, 2, 1, 0, 0, 1, 0, time.UTC), 0, 0)
		}
		if conversion != results[i] {
			t.Errorf("Expected with %s for %f s to match %f, start: %s, end: %s, got %f",
				interval, time_value.Seconds(), results[i], start_t, end_t, conversion)
		}
	}
}

// calculating with given lengths

func TestMonthlyWithLengths(t *testing.T) {
	interval := "MONTHLY"
	results := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 4}
	if len(results) != len(DIF_TIME_VALUES) {
		t.Errorf("Expected results amd TIME_VALUES same length")
	}
	start_t := time.Date(time.Now().Year(), time.Now().Month(), 0, 0, 0, 0, 0, time.UTC)
	var end_t time.Time
	var conversion float64
	for i, time_value := range DIF_TIME_VALUES {
		end_t = start_t.Add(time_value)
		if start_t.Month() != end_t.Month() {
			continue
		} else {
			conversion = numberOfIntervals(end_t.Sub(start_t), interval, time.Now(), 7, 0)
		}
		if conversion != results[i] {
			t.Errorf("[%d]Expected with %s for %f s to match %f, start: %s, end: %s, got %f",
				i, interval, time_value.Seconds(), results[i], start_t, end_t, conversion)
		}
	}
}

func TestYearlyWithLengths(t *testing.T) {
	interval := "MONTHLY"
	results := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	if len(results) != len(DIF_TIME_VALUES) {
		t.Errorf("Expected results amd TIME_VALUES same length")
	}
	start_t := time.Date(time.Now().Year(), time.Now().Month(), 0, 0, 0, 0, 0, time.UTC)
	var end_t time.Time
	var conversion float64
	for i, time_value := range DIF_TIME_VALUES {
		end_t = start_t.Add(time_value)
		if start_t.Month() != end_t.Month() {
			continue
		} else {
			conversion = numberOfIntervals(end_t.Sub(start_t), interval, time.Now(), 0, 365)
		}
		if conversion != results[i] {
			t.Errorf("[%d]Expected with %s for %f s to match %f, start: %s, end: %s, got %f",
				i, interval, time_value.Seconds(), results[i], start_t, end_t, conversion)
		}
	}
}

func TestWithNoInterval(t *testing.T) {
	interval := ""
	start_t := time.Date(time.Now().Year(), time.Now().Month(), 0, 0, 0, 0, 0, time.UTC)
	var end_t time.Time
	var conversion float64
	for i, time_value := range DIF_TIME_VALUES {
		end_t = start_t.Add(time_value)
		if start_t.Month() != end_t.Month() {
			continue
		} else {
			conversion = numberOfIntervals(end_t.Sub(start_t), interval, time.Now(), 0, 365)
		}
		if conversion != 1 {
			t.Errorf("[%d]Expected with %s for %f s to match %f, start: %s, end: %s, got %f",
				i, interval, time_value.Seconds(), 1, start_t, end_t, conversion)
		}
	}
}
