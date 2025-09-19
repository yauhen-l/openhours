package parser

import (
	"bytes"
	"fmt"
	"time"
)

// Matcher represents something that can match against a time
type Matcher interface {
	Match(t time.Time) bool
}

// Parse parses an opening hours string and returns a matcher or errors
func Parse(s string) (Matcher, []error) {
	lex := NewLexer(bytes.NewBufferString(s))
	yyParse(lex)
	switch x := lex.parseResult.(type) {
	case []error:
		return nil, x
	case monthly:
		return x, nil
	default:
		return nil, []error{fmt.Errorf("unsupported result: %T", x)}
	}
}