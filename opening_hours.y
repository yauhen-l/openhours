//go:generate go tool yacc -o opening_hours.go opening_hours.y

%{
package main

import (
//"bufio"
	"fmt"
//	"os"
//"unicode"
)

type TimeRange struct {
	Start int
	End int
}

func NewTimeRange(start, end int) TimeRange {
		return TimeRange{Start: start, End: end}
}

type WeekdayS struct {
	Number int
	Range TimeRange
}

%}

%union {
	num int
	tr TimeRange
	trs []TimeRange
}

%type <num> time hour minute
%type <tr> timespan selector always
%type	<trs> selectors selector_seq

%token <num> NUM WEEKDAY
%token ALWAYS
																								
%left '-' ':'

%start root

%%

root:
				|
								selectors
								{
										yylex.(*Lexer).parseResult = $1
								}
				;

selectors:			selector selector_seq
								{
										$$ = append($2, $1)
								}
				;

selector_seq:  /* empty */
								{ $$ = make([]TimeRange, 0)}
				|
								';' selectors {{ $$ = $2 }}
				;


selector:       always
				|
								WEEKDAY ' ' timespan
								{
										fmt.Printf("Weekday: %d\n", $1)
										$$ = $3
								}
				;

always:         ALWAYS {$$ = NewTimeRange(0,1440)}
				;

timespan:				time '-' time
								{
    								ts := NewTimeRange($1, $3)

    								if ts.Start >= ts.End {
														yylex.Error(fmt.Sprintf("invalid timerange: %v\n", ts))
		    						}
				    				$$ = ts
								}
				;

time:						hour ':' minute
								{
										t := $1 + $3
										if t > 1440 { // > 24:00
														yylex.Error(fmt.Sprintf("invalid time: %d\n", t))
										}
								    $$ = t
								}
				;

hour:						NUM
										{
												if $1 < 0 || $1 > 24 {
																yylex.Error(fmt.Sprintf("invalid hour: %d\n", $1))
												}
												$$ = $1 * 60
										}
				;

minute:					NUM
										{
												if $1 < 0 || $1 > 59 {
																yylex.Error(fmt.Sprintf("invalid minutes: %d\n", $1))
												}
												$$ = $1
										}
				;

%%

								/*
func CompileOpeningHours(s string) ([]TimeRange, err) {
		defer
}
								*/
