package rate


import(
	"time"
	"math"
	"fmt"
)

const (
	MINUTELY  		 	= time.Minute
	HOURLY   			= time.Hour
	DAILY    			= 24 * time.Hour
	WEEKLY   			= 7 * DAILY
)

var   MONTHLY  = time.Duration(daysMonth(time.Now().Year(),time.Now().Month())) * DAILY
var	  YEARLY   = time.Duration(daysYear(time.Now().Year())) * DAILY

func daysMonth(year int,m time.Month) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func daysYear(year int) int {
	return int(time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).
		Sub(time.Date(year, time.January, 0, 0, 0, 0, 0, time.UTC)).Hours()/24)
}


func number_of_intervals(period time.Duration, interval string, calculation_date time.Time, days_in_month int, days_in_year int) float64 {
	if period == 0 {
		return 1
	}
	return calculateTimeSpan(period, interval, calculation_date, days_in_month, days_in_year)
}

func calculatePeriod(period time.Duration, timeSpan float64) float64 {
	var result float64
	if result = 1; math.Mod(period.Seconds(), timeSpan) == 0 {
		result = 0
	}
	return float64(int(period.Seconds()/timeSpan)) + result
}

func calculateTimeSpan(period time.Duration,interval string, calculation_date time.Time, days_in_month int, days_in_year int) float64{
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
			if days_in_month == 0 {
				return calculatePeriod(period, MONTHLY.Seconds())
			}
			fmt.Printf("%d",float64(days_in_month) * DAILY.Seconds())
			return calculatePeriod(period, float64(days_in_month) * DAILY.Seconds())
		case "YEARLY":
			if days_in_year == 0 {
				return calculatePeriod(period, YEARLY.Seconds())
			}
			return calculatePeriod(period, float64(days_in_year) * DAILY.Seconds())
		default:
			return 0
		}

}
