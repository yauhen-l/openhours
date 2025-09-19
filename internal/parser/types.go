package parser

import (
	"time"
)

var any = -1
var wholeWeek = []int{-1}
var wholeDay = []timeRange{newTimeRange(0, 1440)}
var anyTime = appendWeeklyTimeRanges(make(weekly), wholeWeek, wholeDay)

type timeRange struct {
	Start int
	End   int
}

func (tr timeRange) Match(t time.Time) bool {
	minutes := int(t.Hour()*60 + t.Minute())
	return tr.Start <= minutes && minutes < tr.End
}

type monthly map[int]map[int]weekly

func (m monthly) Match(t time.Time) bool {
	for _, month := range []int{any, int(t.Month()) - 1} {
		d, ok := m[month]
		if ok {
			for _, day := range []int{any, int(t.Day())} {
				w, ok := d[day]
				if ok && w.Match(t) {
					return true
				}
			}
		}
	}

	return false
}

func (m monthly) MatchExt(t time.Time) (isOpen bool, nextChange time.Time, duration time.Duration) {
	isOpen = m.Match(t)
	currentMinutes := t.Hour()*60 + t.Minute()

	// Find the applicable weekly schedule for this time
	var applicableWeekly weekly
	found := false

	for _, month := range []int{any, int(t.Month()) - 1} {
		if d, ok := m[month]; ok {
			for _, day := range []int{any, int(t.Day())} {
				if w, ok := d[day]; ok {
					applicableWeekly = w
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}

	if !found {
		// No schedule found, return current status with no change
		return isOpen, time.Time{}, 0
	}

	// Check if this is a "24/7" scenario (always open)
	if applicableWeekly.isAlwaysOpen() {
		return true, time.Time{}, 0
	}

	nextChange = applicableWeekly.findNextChange(t, currentMinutes, isOpen)
	if !nextChange.IsZero() {
		duration = nextChange.Sub(t)
	}

	return isOpen, nextChange, duration
}

type weekly map[int][]timeRange

func (w weekly) Match(t time.Time) bool {
	for _, weekday := range []int{any, int(t.Weekday())} {
		wd, ok := w[weekday]
		if ok {
			for _, tr := range wd {
				if tr.Match(t) {
					return true
				}
			}
		}
	}
	return false
}

func (w weekly) isAlwaysOpen() bool {
	// Check if this weekly schedule equals anyTime (24/7)
	return w.equals(anyTime)
}

func (w weekly) equals(other weekly) bool {
	if len(w) != len(other) {
		return false
	}

	for weekday, ranges := range w {
		otherRanges, exists := other[weekday]
		if !exists || len(ranges) != len(otherRanges) {
			return false
		}

		for i, tr := range ranges {
			if tr.Start != otherRanges[i].Start || tr.End != otherRanges[i].End {
				return false
			}
		}
	}

	return true
}

func (w weekly) findNextChange(t time.Time, currentMinutes int, isCurrentlyOpen bool) time.Time {
	currentWeekday := int(t.Weekday())

	// Check for status changes on the current day first
	if nextChange := w.findNextChangeOnDay(t, currentWeekday, currentMinutes, isCurrentlyOpen); !nextChange.IsZero() {
		return nextChange
	}

	// Check the next 7 days for status changes
	for daysAhead := 1; daysAhead <= 7; daysAhead++ {
		checkDate := t.AddDate(0, 0, daysAhead)
		checkWeekday := int(checkDate.Weekday())

		// Check if this day has any time ranges
		if ranges, exists := w[checkWeekday]; exists && len(ranges) > 0 {
			// If currently open and this day has ranges, we'll close at start of first range
			// If currently closed and this day has ranges, we'll open at start of first range
			firstRange := w.getEarliestRange(checkWeekday)
			if firstRange != nil {
				changeTime := time.Date(checkDate.Year(), checkDate.Month(), checkDate.Day(),
					firstRange.Start/60, firstRange.Start%60, 0, 0, t.Location())
				return changeTime
			}
		} else if ranges, exists := w[any]; exists && len(ranges) > 0 {
			// Check 'any' day rules
			firstRange := w.getEarliestRange(any)
			if firstRange != nil {
				changeTime := time.Date(checkDate.Year(), checkDate.Month(), checkDate.Day(),
					firstRange.Start/60, firstRange.Start%60, 0, 0, t.Location())
				return changeTime
			}
		}
	}

	// No change found in the next week
	return time.Time{}
}

func (w weekly) findNextChangeOnDay(t time.Time, weekday, currentMinutes int, isCurrentlyOpen bool) time.Time {
	allRanges := []timeRange{}

	// Collect all relevant time ranges for this weekday
	if ranges, exists := w[weekday]; exists {
		allRanges = append(allRanges, ranges...)
	}
	if ranges, exists := w[any]; exists {
		allRanges = append(allRanges, ranges...)
	}

	if len(allRanges) == 0 {
		return time.Time{}
	}

	// Sort ranges by start time
	for i := 0; i < len(allRanges)-1; i++ {
		for j := i + 1; j < len(allRanges); j++ {
			if allRanges[i].Start > allRanges[j].Start {
				allRanges[i], allRanges[j] = allRanges[j], allRanges[i]
			}
		}
	}

	// Find the next status change
	if isCurrentlyOpen {
		// Looking for closing time - need to find the latest end time of all overlapping ranges
		latestEnd := -1
		inAnyRange := false

		for _, tr := range allRanges {
			if currentMinutes >= tr.Start && currentMinutes < tr.End {
				inAnyRange = true
				if tr.End > latestEnd {
					latestEnd = tr.End
				}
			}
		}

		if inAnyRange && latestEnd > currentMinutes {
			return time.Date(t.Year(), t.Month(), t.Day(), latestEnd/60, latestEnd%60, 0, 0, t.Location())
		}
	} else {
		// Looking for opening time - start of next range
		for _, tr := range allRanges {
			if currentMinutes < tr.Start {
				return time.Date(t.Year(), t.Month(), t.Day(), tr.Start/60, tr.Start%60, 0, 0, t.Location())
			}
		}
	}

	return time.Time{}
}

func (w weekly) getEarliestRange(weekday int) *timeRange {
	ranges, exists := w[weekday]
	if !exists || len(ranges) == 0 {
		return nil
	}

	earliest := &ranges[0]
	for i := 1; i < len(ranges); i++ {
		if ranges[i].Start < earliest.Start {
			earliest = &ranges[i]
		}
	}
	return earliest
}

func makeMonthly(month int, days []int) monthly {
	m := make(monthly)
	m[month] = make(map[int]weekly)
	for _, day := range days {
		m[month][day] = make(weekly)
	}
	return m
}

func setWeekly(monthly monthly, weekly weekly) monthly {
	for _, days := range monthly {
		for d, _ := range days {
			days[d] = weekly
		}
	}
	return monthly
}

func mergeMonthly(m1, m2 monthly) monthly {
	for m, _ := range m2 {
		if _, ok := m1[m]; ok {
			m1[m] = mergeMonthdays(m1[m], m2[m])
		} else {
			m1[m] = m2[m]
		}
	}
	return m1
}

func mergeMonthdays(d1, d2 map[int]weekly) map[int]weekly {
	for d, _ := range d2 {
		if _, ok := d1[d]; ok {
			d1[d] = mergeWeeklyTimeRanges(d1[d], d2[d])
		} else {
			d1[d] = d2[d]
		}
	}
	return d1
}

func newTimeRange(start, end int) timeRange {
	return timeRange{Start: start, End: end}
}

func appendWeeklyTimeRanges(weekly weekly, weeks []int, trs []timeRange) weekly {
	for _, w := range weeks {
		weekly[w] = append(weekly[w], trs...)
	}
	return weekly
}

func mergeWeeklyTimeRanges(w1, w2 weekly) weekly {
	for i, _ := range w2 {
		if _, ok := w1[i]; ok {
			w1[i] = append(w1[i], w2[i]...)
		} else {
			w1[i] = w2[i]
		}
	}
	return w1
}
