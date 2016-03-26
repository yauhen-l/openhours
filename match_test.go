package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func sunday() time.Time {
	return time.Date(2016, time.March, 20, 0, 0, 0, 0, time.UTC)
}

func TestMatch(t *testing.T) {
	assert := assert.New(t)

	for wd, i := range weekdays {
		t := sunday().AddDate(0, 0, i)
		fmt.Println(t)

		interval, errs := CompileOpeningHours(wd)
		assert.Empty(errs)

		fmt.Println(interval)

		assert.True(interval.Match(t))
		assert.False(interval.Match(t.Add(-1 * time.Minute)))

		assert.False(interval.Match(t.AddDate(0, 0, 1)))
		assert.True(interval.Match(t.AddDate(0, 0, 1).Add(-1 * time.Minute)))
	}
}
