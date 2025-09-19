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
	matcher    parser.Matcher
	definition string
}

// Match returns true if provided time matches current OpenHours
func (oh *OpenHours) Match(t time.Time) bool {
	return oh.matcher.Match(t)
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
