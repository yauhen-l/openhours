//go:generate go tool yacc -o opening_hours.go opening_hours.y

%{
package main

import (
	"fmt"
  "time"
  "bytes"
)

var any = -1
var wholeWeek = []int{-1}
var wholeDay = []TimeRange{NewTimeRange(0,1440)}

type OpeningHours struct {
		data       Monthly // Month -> Day -> Weekday -> Hours
		definition string
}

func (oh *OpeningHours) Match(t time.Time) bool {
		return oh.data.Match(t)
}

func (oh *OpeningHours) Definition() string {
		return oh.definition
}

type TimeRange struct {
	Start int
	End int
}

func (tr TimeRange) Match(t time.Time) bool {
    minutes := int(t.Hour() * 60 + t.Minute())
		return tr.Start <= minutes && minutes < tr.End
}

type Monthly map[int]map[int]Weekly

func (m Monthly) Match(t time.Time) bool {
		for _, month := range []int{any, int(t.Month()) - 1 } {
			d, ok := m[month]
			if ok {
				for _, day := range []int{any, int(t.Day()) - 1 } {
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

%}

%union {
	num int
	nums []int
	tr TimeRange
	trs []TimeRange
	weekly Weekly
	monthly Monthly
}

%type		<num>						time hour minute day DIGIT NUM
%type		<nums>					ws weekdays weekdays_seq days ds more_days
%type		<tr>						timespan
%type		<trs>						timespans timespans_seq
%type		<monthly>				selector selectors selector_seq
%type		<weekly>				small_range_selector
												
%token <num> WEEKDAY MONTH
%token ALWAYS
%token SEP

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

selectors:			selector selector_seq { $$ = mergeMonthly($1, $2) }
				;

selector_seq:  /* empty */
								{ $$ = make(Monthly) }
				|
								';' selectors { $$ = $2 }
				;


selector:       ALWAYS
								{
										$$ = setWeekly(makeMonthly(any, []int{ any }),
																	 appendWeeklyTimeRanges(make(Weekly),
																													wholeWeek,
																													wholeDay))
								}
				|
								days SEP MONTH SEP small_range_selector
								{
								    $$ = setWeekly(makeMonthly($3, $1), $5)
								}
				|
								days SEP small_range_selector
								{
								    $$ = setWeekly(makeMonthly(any, $1), $3)
								}
				|
								MONTH SEP small_range_selector
								{
								    $$ = setWeekly(makeMonthly($1, []int{ any }), $3)
								}
				|								
								small_range_selector
								{
										$$ = setWeekly(makeMonthly(any, []int{ any }), $1)
								}
				;


days:						ds more_days
								{
										$$ = append($1, $2...)
								}
				;

ds:							day
								{ $$ = []int{ $1 } }
				|
								day '-' day
								{
										if $1 > $3 {
														yylex.Error(fmt.Sprintf("invalid days range: %d - %d\n", $1, $3))
										}
								    $$ = make([]int, 0)
								    for i := $1; i <= $3; i++ {
												$$ = append($$, i)
										}
								}
				;

more_days:			/* empty */
								{ $$ = []int{} }
				|
								',' days
								{ $$ = $2 }
				;

small_range_selector:
								weekdays
								{$$ = appendWeeklyTimeRanges(make(Weekly), $1, wholeDay)}
				|				weekdays SEP timespans
								{$$ = appendWeeklyTimeRanges(make(Weekly), $1, $3)}
				|				timespans
								{$$ = appendWeeklyTimeRanges(make(Weekly), wholeWeek, $1)}
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

day:						NUM
								{
										if $1 < 1 || $1 > 31 { yylex.Error(fmt.Sprintf("invalid day: %d\n", $1)) }
										$$ = $1
								}
				;

NUM:						DIGIT DIGIT
								{$$ = $1 * 10 + $2}
				;

DIGIT:					'0' {$$=0} | '1' {$$=1} | '2' {$$=2} | '3' {$$=3} | '4' {$$=4} | '5' {$$=5}	| '6' {$$=6} | '7' {$$=7} | '8' {$$=8} | '9' {$$=9}
				;

%%


func CompileOpeningHours(s string) (*OpeningHours, []error) {
	lex := NewLexer(bytes.NewBufferString(s))
	yyParse(lex)
	switch x := lex.parseResult.(type){
	case []error:
	  return nil, x
	case Monthly:
  	return &OpeningHours{data: x, definition: s}, nil
	default:
 	  return nil, []error{fmt.Errorf("unsupported result: %T", x)}
	}
}

