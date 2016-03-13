//go:generate go tool yacc -o opening_hours.go opening_hours.y

%{
package main

import (
	"fmt"
)

var wholeWeek = []int{0,1,2,3,4,5,6}

type TimeRange struct {
	Start int
	End int
}

func NewTimeRange(start, end int) TimeRange {
		return TimeRange{Start: start, End: end}
}

func makeWeekly() [][]TimeRange {
  weekly := make([][]TimeRange, 7)
	for i, _:= range weekly {
    weekly[i] = make([]TimeRange, 0)
	}
  return weekly
}

func appendWeeklyTimeRanges(weekly [][]TimeRange, weeks []int, trs []TimeRange) [][]TimeRange {
  for _, w := range weeks {
  	weekly[w] = append(weekly[w], trs...)
	}
  return weekly
}

func mergeWeeklyTimeRanges(w1, w2 [][]TimeRange) [][]TimeRange {
	for i, _ := range w1 {
    w1[i] = append(w1[i], w2[i]...)
	}
  return w1
}

%}

%union {
	num int
	nums []int
	continuous bool
	tr TimeRange
	trs []TimeRange
	weekly [][]TimeRange
}

%type		<num>						time hour minute DIGIT NUM
%type		<nums>					ws weekdays weekdays_seq
%type		<tr>						timespan
%type		<trs>						timespans timespans_seq always
%type		<weekly>				selector selectors selector_seq
												
%token <num> WEEKDAY
%token ALWAYS
																								
%left '-' ':'

%start root

%%

root:
				|
								selectors
								{
										res := yylex.(*Lexer).parseResult
										if res == nil {
														yylex.(*Lexer).parseResult = $1
										}
								}
				;

selectors:			selector selector_seq { $$ = mergeWeeklyTimeRanges($1, $2) }
				;

selector_seq:  /* empty */
									 { $$ = makeWeekly() }
				|
								';' selectors {{ $$ = $2 }}
				;


selector:       always {$$ = appendWeeklyTimeRanges(makeWeekly(), wholeWeek, $1)}
				|
								weekdays ' ' timespans
								{
										$$ = appendWeeklyTimeRanges(makeWeekly(), $1, $3)
								}
				|				timespans {$$ = appendWeeklyTimeRanges(makeWeekly(), wholeWeek, $1)}
				;

always:         ALWAYS {$$ = []TimeRange{NewTimeRange(0,1440)}}
				;

weekdays:       ws	weekdays_seq
								{
										$$ = append($1, $2...)
								}
				;

ws:							WEEKDAY
								{
										$$ = []int{$1}
								}
				|
								WEEKDAY '-' WEEKDAY
								{
										if $1 > $3 {
														yylex.Error(fmt.Sprintf("invalid weekdays range: %d - %d\n", $1, $3))
										}
								    $$ = make([]int, 0)
								    for i := $1; i <= $3; i++ {
												$$ = append($$, i)
										}
								}
				;

weekdays_seq:		/* empty */
								{
										$$ = []int{}
								}
				|
								',' weekdays
								{
										$$ = $2
								}
				;

timespans:			timespan timespans_seq	{ $$ = append($2, $1) }
				;

timespans_seq:	/* empty */   { $$ = []TimeRange{} }
				|
								',' timespans { $$ = $2 }
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

NUM:						DIGIT DIGIT
								{$$ = $1 * 10 + $2}
				;

DIGIT:					'0' {$$=0} | '1' {$$=1} | '2' {$$=2}
								| '3' {$$=3} | '4' {$$=4} | '5' {$$=5}
								| '6' {$$=6} | '7' {$$=7} | '8' {$$=8} | '9' {$$=9}
				;

%%

								/*
func CompileOpeningHours(s string) ([]TimeRange, err) {
		defer
}
								*/
