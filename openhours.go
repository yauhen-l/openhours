// Package oenhours provides functions for parsing time intervals
// defined in human-readable form and check whether specific time matches these intervals

//go:generate nex -o openhours_lexer.go -e openhours.nex
//go:generate go tool yacc -o openhours_parser.go openhours.y
package openhours

import (
	"bytes"
	"fmt"
	"time"
)

var any = -1
var wholeWeek = []int{-1}
var wholeDay = []TimeRange{NewTimeRange(0, 1440)}
var anyTime = appendWeeklyTimeRanges(make(Weekly), wholeWeek, wholeDay)

// OpenHours is a main top-level struct to work with
type OpenHours struct {
	data       Monthly // Month -> Day -> Weekday -> Hours
	definition string
}

// Match returns true if provided time matches current OpenHours
func (oh *OpenHours) Match(t time.Time) bool {
	return oh.data.Match(t)
}

// Definition returns initial open hours string
func (oh *OpenHours) Definition() string {
	return oh.definition
}

type TimeRange struct {
	Start int
	End   int
}

func (tr TimeRange) Match(t time.Time) bool {
	minutes := int(t.Hour()*60 + t.Minute())
	return tr.Start <= minutes && minutes < tr.End
}

type Monthly map[int]map[int]Weekly

func (m Monthly) Match(t time.Time) bool {
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

type Weekly map[int][]TimeRange

func (w Weekly) Match(t time.Time) bool {
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

func makeMonthly(month int, days []int) Monthly {
	m := make(Monthly)
	m[month] = make(map[int]Weekly)
	for _, day := range days {
		m[month][day] = make(Weekly)
	}
	return m
}

func setWeekly(monthly Monthly, weekly Weekly) Monthly {
	for _, days := range monthly {
		for d, _ := range days {
			days[d] = weekly
		}
	}
	return monthly
}

func mergeMonthly(m1, m2 Monthly) Monthly {
	for m, _ := range m2 {
		if _, ok := m1[m]; ok {
			m1[m] = mergeMonthdays(m1[m], m2[m])
		} else {
			m1[m] = m2[m]
		}
	}
	return m1
}

func mergeMonthdays(d1, d2 map[int]Weekly) map[int]Weekly {
	for d, _ := range d2 {
		if _, ok := d1[d]; ok {
			d1[d] = mergeWeeklyTimeRanges(d1[d], d2[d])
		} else {
			d1[d] = d2[d]
		}
	}
	return d1
}

func NewTimeRange(start, end int) TimeRange {
	return TimeRange{Start: start, End: end}
}

func appendWeeklyTimeRanges(weekly Weekly, weeks []int, trs []TimeRange) Weekly {
	for _, w := range weeks {
		weekly[w] = append(weekly[w], trs...)
	}
	return weekly
}

func mergeWeeklyTimeRanges(w1, w2 Weekly) Weekly {
	for i, _ := range w2 {
		if _, ok := w1[i]; ok {
			w1[i] = append(w1[i], w2[i]...)
		} else {
			w1[i] = w2[i]
		}
	}
	return w1
}

// CompileOpenHours parses open hours string
// It returns OpenHours object or list of compilation errors
func CompileOpenHours(s string) (*OpenHours, []error) {
	lex := NewLexer(bytes.NewBufferString(s))
	yyParse(lex)
	switch x := lex.parseResult.(type) {
	case []error:
		return nil, x
	case Monthly:
		return &OpenHours{data: x, definition: s}, nil
	default:
		return nil, []error{fmt.Errorf("unsupported result: %T", x)}
	}
}
