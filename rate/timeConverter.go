package rate

import (
	"math"
	"time"
)

/*

Report for Minute, Hour, Day and Week
*/
const (
	MINUTELY = time.Minute    // Report set by minute
	HOURLY   = time.Hour      // Report set by hour
	DAILY    = 24 * time.Hour // Report set by day
	WEEKLY   = 7 * DAILY      // Report set by week
)

// MONTHLY Report set by month
var MONTHLY = time.Duration(daysMonth(time.Now().Year(), time.Now().Month())) * DAILY

// YEARLY Report set by year
var YEARLY = time.Duration(daysYear(time.Now().Year())) * DAILY

func daysMonth(year int, m time.Month) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func daysYear(year int) int {
	return int(time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).
		Sub(time.Date(year, time.January, 0, 0, 0, 0, 0, time.UTC)).Hours() / 24)
}

func numberOfIntervals(period time.Duration, interval string, calculationDate time.Time, daysMonth int, daysYear int) float64 {
	if period == 0 {
		return 1
	}
	return calculateTimeSpan(period, interval, calculationDate, daysMonth, daysYear)
}

func calculatePeriod(period time.Duration, timeSpan float64) float64 {
	var result float64
	if result = 1; math.Mod(period.Seconds(), timeSpan) == 0 {
		result = 0
	}
	return float64(int(period.Seconds()/timeSpan)) + result
}

func calculateTimeSpan(period time.Duration, interval string, calculationDate time.Time, daysMonth int, daysYear int) float64 {
	switch interval {
	case "MINUTELY":
		return calculatePeriod(period, MINUTELY.Seconds())
	case "HOURLY":
		return calculatePeriod(period, HOURLY.Seconds())
	case "DAILY":
		return calculatePeriod(period, DAILY.Seconds())
	case "WEEKLY":
		return calculatePeriod(period, WEEKLY.Seconds())
	case "MONTHLY":
		if daysMonth == 0 {
			return calculatePeriod(period, MONTHLY.Seconds())
		}
		return calculatePeriod(period, float64(daysMonth)*DAILY.Seconds())
	case "YEARLY":
		if daysYear == 0 {
			return calculatePeriod(period, YEARLY.Seconds())
		}
		return calculatePeriod(period, float64(daysYear)*DAILY.Seconds())
	default:
		return 1
	}

}
