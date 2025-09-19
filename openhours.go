// Package openhours provides functions for parsing time intervals
// defined in human-readable form and check whether specific time matches these intervals

//go:generate nex -o internal/parser/openhours_lexer.go -e internal/parser/openhours.nex
//go:generate goyacc -o internal/parser/openhours_parser.go internal/parser/openhours.y
package openhours

import (
	"time"

	"github.com/yauhen-l/openhours/internal/parser"
)

// OpenHours is a main top-level struct to work with
type OpenHours struct {
	matcher    parser.ExtendedMatcher
	definition string
}

// Match returns true if provided time matches current OpenHours
func (oh *OpenHours) Match(t time.Time) bool {
	return oh.matcher.Match(t)
}

// MatchExt returns detailed information about opening hours status:
// - isOpen: whether the location is currently open at the given time
// - nextChange: when the status will next change (open->closed or closed->open)
// - duration: how long until the status change occurs
// If no change is found (e.g., for "24/7"), nextChange will be zero time and duration will be 0.
func (oh *OpenHours) MatchExt(t time.Time) (isOpen bool, nextChange time.Time, duration time.Duration) {
	return oh.matcher.MatchExt(t)
}

// Definition returns initial open hours string
func (oh *OpenHours) Definition() string {
	return oh.definition
}

// CompileOpenHours parses open hours string
// It returns OpenHours object or list of compilation errors
func CompileOpenHours(s string) (*OpenHours, []error) {
	matcher, errs := parser.Parse(s)
	if errs != nil {
		return nil, errs
	}
	return &OpenHours{matcher: matcher, definition: s}, nil
}
