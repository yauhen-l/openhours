# openhours
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
