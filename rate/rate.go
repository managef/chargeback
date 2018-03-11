package rate

import (
	"time"
)

/*
Format Rate per time, value and rate presentation
*/
type Format struct {
	Time  string
	Value float64
}

/*
Tier of a rate with Fixed and variable cost
*/
type Tier struct {
	Fixed    Format
	Variable Format
}

func occurrence(tier Tier, value float64, cycleDuration time.Duration, date time.Time) float64 {
	// Returns fixed_cost always + variable_cost sometimes
	// Fixed cost are always added fully, variable costs are only added if value is not nil
	// fix_inter: number of intervals in the calculation => how many times do we need to apply the rate to get a monthly (cycle) rate (min = 1)
	// fix_inter * fixed_rate ==  interval_rate (i.e. monthly)
	// var_inter * variable_rate == interval_rate (i.e. monthly)
	fixInter := numberOfIntervals(cycleDuration, tier.Fixed.Time, date, 0, 0)
	varInter := numberOfIntervals(cycleDuration, tier.Variable.Time, date, 0, 0)
	var extraValue float64
	if extraValue = 0; value != 0 {
		extraValue = varInter * tier.Variable.Value
	}
	return fixInter*varInter + extraValue
}
