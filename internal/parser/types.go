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