# openhours [![Build Status](https://travis-ci.org/yauhen-l/openhours.png?branch=master)](https://travis-ci.org/yauhen-l/openhours)
golang implementation of opening hours (inspired by OpenStreetMap opening hours: http://wiki.openstreetmap.org/wiki/Key:opening_hours)

This Go library let you define time intervals in human-readable form and then check if specified time matches this interval.

This library uses [dep](https://github.com/golang/dep) to manage dependencies.

## Install

Run `go get github.com/yauhen-l/openhours`

Supports Golang 1.7+ (due to `t.Run` usages in tests, but should work for earlier versions too)

## Usage

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/yauhen-l/openhours"
)

func main() {
	oh, errs := openhours.CompileOpenHours("Mo-We 06:00-17:00")
	if len(errs) > 0 {
	  fmt.Printf("%v\n", errs)
	  return
	}
	fmt.Println(oh.Match(time.Now()))
}
```
## Examples
OpenHours simplified pattern:
```
OpenHours = (days(,days)*)? (MMM)? (Wds(,Wds)*)? (timespan(,timespans)*)?

days = day | day-day       #Days of month
Wds = Wd | Wd-Wd           #Weekdays
timespan = hh:mm-hh:mm     #Day time range
```
Openhours           |Description
--------------------|-----------
24/7                |Matches everything
24 Dec              |Any time on 24th of December
01-05               |Any time from 1st till 5th of any month
18:00-18:30         |Any day from 6PM till 6:30PM
Sa,Su 10:00-22:00   |From 10AM till 10PM on Saturday and Sunday
Mo; Tu 9:00-15:00   |Any time on Monday and from 9AM till 3PM on Tuesday

## Contribution

If want to extend syntax then you need to install few more packages:
```
// Lexer
go get github.com/blynn/nex
// Parser
go get golang.org/x/tools/cmd/goyacc
```

To verify changes:
```
go generate && go test -v
```