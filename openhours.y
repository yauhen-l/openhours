%{
package openhours

import "fmt"
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
										$$ = setWeekly(makeMonthly(any, []int{ any }), anyTime)
								}
				|
								days
								{
										$$ = setWeekly(makeMonthly(any, $1), anyTime)
								}
				|
								days SEP MONTH
								{
								    $$ = setWeekly(makeMonthly($3, $1), anyTime)
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
