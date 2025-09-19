# openhours [![Build Status](https://github.com/yauhen-l/openhours/workflows/CI/badge.svg)](https://github.com/yauhen-l/openhours/actions)
golang implementation of opening hours (inspired by OpenStreetMap opening hours: http://wiki.openstreetmap.org/wiki/Key:opening_hours)

This Go library let you define time intervals in human-readable form and then check if specified time matches this interval.

## Install

Run `go get github.com/yauhen-l/openhours`

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

## OpenStreetMap Specification Compatibility

This table shows which features from the [OpenStreetMap opening_hours specification](https://wiki.openstreetmap.org/wiki/Key:opening_hours) are currently supported:

| Feature | Status | Examples | Notes |
|---------|--------|----------|-------|
| **Time Formats** |
| Basic 24-hour (HH:MM) | ✅ | `08:00-17:00` | Fully supported |
| Multiple time intervals | ✅ | `08:00-12:00,13:00-17:00` | Comma-separated |
| Time validation | ✅ | Hours 0-24, minutes 0-59 | Full validation |
| 24:00+ format | ❌ | `18:00-26:00` | Not supported |
| Open-ended times | ❌ | `18:00+` | Not supported |
| Variable times | ❌ | `sunrise-sunset` | Not supported |
| **Weekdays** |
| Individual weekdays | ✅ | `Mo`, `Tu`, `We`, `Th`, `Fr`, `Sa`, `Su` | All supported |
| Weekday ranges | ✅ | `Mo-Fr`, `Sa-Su` | Fully supported |
| Multiple weekdays | ✅ | `Mo,We,Fr` | Comma-separated |
| Nth occurrence | ❌ | `Su[1]` (first Sunday) | Not supported |
| **Dates and Months** |
| Day of month | ✅ | `20` (20th of any month) | Fully supported |
| Day ranges | ✅ | `01-05`, `20-25` | Fully supported |
| Month names | ✅ | `Jan`, `Feb`, `Mar`, etc. | All abbreviations |
| Day with month | ✅ | `20 Mar`, `24 Dec` | Fully supported |
| Specific years | ❌ | `2024 Jan 10` | Not supported |
| Week numbers | ❌ | `week 25` | Not supported |
| **Special Values** |
| Always open | ✅ | `24/7` | Fully supported |
| Public holidays | ❌ | `PH`, `PH off` | Not supported |
| School holidays | ❌ | `SH` | Not supported |
| Closed/off status | ❌ | `off`, `closed` | Not supported |
| **Complex Rules** |
| Semicolon separator | ✅ | `Mo 10:00-12:00; Tu 14:00-16:00` | Overwriting rules |
| Combined patterns | ✅ | `20,21 Mar Mo` | Day+month+weekday |
| Comments | ❌ | `"children only"` | Not supported |
| Colon separator | ❌ | `Dec 25: 08:30-20:00` | Not supported |
| Fallback rules | ❌ | `Mo-Fr 08:00-18:00 \|\| "by appointment"` | Not supported |

**Coverage Summary:**
- ✅ **Basic features**: ~70% - Covers most common use cases
- ❌ **Advanced features**: ~20% - Missing complex rules and special cases

This implementation focuses on the most commonly used opening hours patterns and provides a solid foundation for typical business hour specifications.

## Contribution

If want to extend syntax then you need to install few more packages:
```
// Lexer
go install github.com/blynn/nex
// Parser
go install github.com/cznic/goyacc
```
And make sure that both `nex` and `goyacc` are in the path


To verify changes:
```
go generate && go test -v
```
