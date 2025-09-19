package openhours

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMatchExt(t *testing.T) {
	t.Run("Business hours Mo-Fr 09:00-17:00", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Mo-Fr 09:00-17:00", assert)

		// Test during business hours (Tuesday 14:00)
		tuesday2pm := time.Date(2016, 3, 22, 14, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(tuesday2pm)

		assert.True(isOpen, "Should be open during business hours")
		assert.False(nextChange.IsZero(), "Should have a next change time")
		assert.Equal(17, nextChange.Hour(), "Should close at 17:00")
		assert.Equal(0, nextChange.Minute(), "Should close at 17:00")
		assert.Equal(3*time.Hour, duration, "Should close in 3 hours")

		// Test after business hours (Tuesday 19:00)
		tuesday7pm := time.Date(2016, 3, 22, 19, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration = oh.MatchExt(tuesday7pm)

		assert.False(isOpen, "Should be closed after business hours")
		assert.False(nextChange.IsZero(), "Should have a next change time")
		assert.Equal(time.Wednesday, nextChange.Weekday(), "Should open next Wednesday")
		assert.Equal(9, nextChange.Hour(), "Should open at 09:00")
		assert.Equal(0, nextChange.Minute(), "Should open at 09:00")
		assert.Equal(14*time.Hour, duration, "Should open in 14 hours")
	})

	t.Run("24/7 always open", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("24/7", assert)

		// Test any time
		anytime := time.Date(2016, 3, 22, 14, 30, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(anytime)

		assert.True(isOpen, "Should always be open")
		assert.True(nextChange.IsZero(), "Should have no next change for 24/7")
		assert.Equal(time.Duration(0), duration, "Should have zero duration for 24/7")
	})

	t.Run("Weekend only Sa,Su 10:00-22:00", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Sa,Su 10:00-22:00", assert)

		// Test Friday evening (closed)
		friday8pm := time.Date(2016, 3, 25, 20, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(friday8pm)

		assert.False(isOpen, "Should be closed on Friday")
		assert.False(nextChange.IsZero(), "Should have a next change time")
		assert.Equal(time.Saturday, nextChange.Weekday(), "Should open on Saturday")
		assert.Equal(10, nextChange.Hour(), "Should open at 10:00")
		assert.Equal(0, nextChange.Minute(), "Should open at 10:00")
		assert.Equal(14*time.Hour, duration, "Should open in 14 hours")

		// Test Saturday afternoon (open)
		saturday2pm := time.Date(2016, 3, 26, 14, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration = oh.MatchExt(saturday2pm)

		assert.True(isOpen, "Should be open on Saturday afternoon")
		assert.False(nextChange.IsZero(), "Should have a next change time")
		assert.Equal(time.Saturday, nextChange.Weekday(), "Should close later on Saturday")
		assert.Equal(22, nextChange.Hour(), "Should close at 22:00")
		assert.Equal(0, nextChange.Minute(), "Should close at 22:00")
		assert.Equal(8*time.Hour, duration, "Should close in 8 hours")
	})

	t.Run("Exact opening time", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Mo-Fr 09:00-17:00", assert)

		// Test exactly at opening time (9:00 AM on Monday)
		mondayExact9am := time.Date(2016, 3, 21, 9, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(mondayExact9am)

		assert.True(isOpen, "Should be open exactly at opening time")
		assert.False(nextChange.IsZero(), "Should have a next change time")
		assert.Equal(17, nextChange.Hour(), "Should close at 17:00")
		assert.Equal(0, nextChange.Minute(), "Should close at 17:00")
		assert.Equal(8*time.Hour, duration, "Should close in 8 hours")
	})

	t.Run("Exact closing time", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Mo-Fr 09:00-17:00", assert)

		// Test exactly at closing time (5:00 PM on Monday)
		mondayExact5pm := time.Date(2016, 3, 21, 17, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(mondayExact5pm)

		assert.False(isOpen, "Should be closed exactly at closing time")
		assert.False(nextChange.IsZero(), "Should have a next change time")
		assert.Equal(time.Tuesday, nextChange.Weekday(), "Should open next Tuesday")
		assert.Equal(9, nextChange.Hour(), "Should open at 09:00")
		assert.Equal(0, nextChange.Minute(), "Should open at 09:00")
		assert.Equal(16*time.Hour, duration, "Should open in 16 hours")
	})

	t.Run("One minute before opening", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Mo-Fr 09:00-17:00", assert)

		// Test one minute before opening (8:59 AM on Monday)
		monday859am := time.Date(2016, 3, 21, 8, 59, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(monday859am)

		assert.False(isOpen, "Should be closed one minute before opening")
		assert.False(nextChange.IsZero(), "Should have a next change time")
		assert.Equal(9, nextChange.Hour(), "Should open at 09:00")
		assert.Equal(0, nextChange.Minute(), "Should open at 09:00")
		assert.Equal(1*time.Minute, duration, "Should open in 1 minute")
	})

	t.Run("One minute before closing", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Mo-Fr 09:00-17:00", assert)

		// Test one minute before closing (4:59 PM on Monday)
		monday459pm := time.Date(2016, 3, 21, 16, 59, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(monday459pm)

		assert.True(isOpen, "Should be open one minute before closing")
		assert.False(nextChange.IsZero(), "Should have a next change time")
		assert.Equal(17, nextChange.Hour(), "Should close at 17:00")
		assert.Equal(0, nextChange.Minute(), "Should close at 17:00")
		assert.Equal(1*time.Minute, duration, "Should close in 1 minute")
	})

	t.Run("Multiple time ranges same day", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Mo 09:00-12:00,14:00-18:00", assert)

		// Test during morning hours
		monday10am := time.Date(2016, 3, 21, 10, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(monday10am)

		assert.True(isOpen, "Should be open during morning hours")
		assert.Equal(12, nextChange.Hour(), "Should close at 12:00 for lunch")
		assert.Equal(2*time.Hour, duration, "Should close in 2 hours")

		// Test during lunch break
		monday1pm := time.Date(2016, 3, 21, 13, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration = oh.MatchExt(monday1pm)

		assert.False(isOpen, "Should be closed during lunch break")
		assert.Equal(14, nextChange.Hour(), "Should reopen at 14:00")
		assert.Equal(1*time.Hour, duration, "Should reopen in 1 hour")

		// Test during afternoon hours
		monday3pm := time.Date(2016, 3, 21, 15, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration = oh.MatchExt(monday3pm)

		assert.True(isOpen, "Should be open during afternoon hours")
		assert.Equal(18, nextChange.Hour(), "Should close at 18:00")
		assert.Equal(3*time.Hour, duration, "Should close in 3 hours")
	})

	t.Run("Weekend transition", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Mo-Fr 09:00-17:00", assert)

		// Test Friday at closing time
		friday5pm := time.Date(2016, 3, 25, 17, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(friday5pm)

		assert.False(isOpen, "Should be closed at Friday 17:00")
		assert.Equal(time.Monday, nextChange.Weekday(), "Should reopen on Monday")
		assert.Equal(9, nextChange.Hour(), "Should reopen at 09:00")
		assert.Equal(64*time.Hour, duration, "Should reopen after weekend (64 hours)")

		// Test Sunday evening
		sunday8pm := time.Date(2016, 3, 27, 20, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration = oh.MatchExt(sunday8pm)

		assert.False(isOpen, "Should be closed on Sunday")
		assert.Equal(time.Monday, nextChange.Weekday(), "Should open on Monday")
		assert.Equal(9, nextChange.Hour(), "Should open at 09:00")
		assert.Equal(13*time.Hour, duration, "Should open in 13 hours")
	})

	t.Run("Single day schedule", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Mo 10:00-16:00", assert)

		// Test on the scheduled day
		monday2pm := time.Date(2016, 3, 21, 14, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(monday2pm)

		assert.True(isOpen, "Should be open on Monday")
		assert.Equal(16, nextChange.Hour(), "Should close at 16:00")
		assert.Equal(2*time.Hour, duration, "Should close in 2 hours")

		// Test on a different day
		tuesday2pm := time.Date(2016, 3, 22, 14, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration = oh.MatchExt(tuesday2pm)

		assert.False(isOpen, "Should be closed on Tuesday")
		assert.Equal(time.Monday, nextChange.Weekday(), "Should open next Monday")
		assert.Equal(10, nextChange.Hour(), "Should open at 10:00")
		// Tuesday 14:00 to next Monday 10:00 = 140 hours (5 days and 20 hours)
		assert.Equal(140*time.Hour, duration, "Should open in 140 hours (next Monday)")
	})

	t.Run("Midnight crossing", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Mo 22:00-23:59", assert)

		// Test near end of day
		monday11pm := time.Date(2016, 3, 21, 23, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(monday11pm)

		assert.True(isOpen, "Should be open at 23:00")
		assert.Equal(23, nextChange.Hour(), "Should close at 23:59")
		assert.Equal(59, nextChange.Minute(), "Should close at 23:59")
		assert.Equal(59*time.Minute, duration, "Should close in 59 minutes")
	})

	t.Run("No schedule (never open)", func(t *testing.T) {
		assert := require.New(t)
		// This should create an empty schedule that's never open
		oh, errs := CompileOpenHours("")
		assert.Error(errs[0], "Empty schedule should produce errors")
		assert.Nil(oh, "Should not create OpenHours for empty schedule")
	})

	t.Run("Early morning edge case", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Mo-Fr 09:00-17:00", assert)

		// Test very early morning (3 AM)
		monday3am := time.Date(2016, 3, 21, 3, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(monday3am)

		assert.False(isOpen, "Should be closed at 3 AM")
		assert.Equal(9, nextChange.Hour(), "Should open at 09:00")
		assert.Equal(6*time.Hour, duration, "Should open in 6 hours")
	})

	t.Run("Late night edge case", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Mo-Fr 09:00-17:00", assert)

		// Test late night (11 PM)
		monday11pm := time.Date(2016, 3, 21, 23, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(monday11pm)

		assert.False(isOpen, "Should be closed at 11 PM")
		assert.Equal(time.Tuesday, nextChange.Weekday(), "Should open next Tuesday")
		assert.Equal(9, nextChange.Hour(), "Should open at 09:00")
		assert.Equal(10*time.Hour, duration, "Should open in 10 hours")
	})

	t.Run("Complex schedule with gaps", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Mo 08:00-10:00,12:00-14:00,16:00-18:00", assert)

		// Test in first gap (10:30 AM)
		monday1030am := time.Date(2016, 3, 21, 10, 30, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(monday1030am)

		assert.False(isOpen, "Should be closed in first gap")
		assert.Equal(12, nextChange.Hour(), "Should reopen at 12:00")
		assert.Equal(90*time.Minute, duration, "Should reopen in 90 minutes")

		// Test in second gap (15:00)
		monday3pm := time.Date(2016, 3, 21, 15, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration = oh.MatchExt(monday3pm)

		assert.False(isOpen, "Should be closed in second gap")
		assert.Equal(16, nextChange.Hour(), "Should reopen at 16:00")
		assert.Equal(1*time.Hour, duration, "Should reopen in 1 hour")
	})

	t.Run("Very short time window", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Mo 12:00-12:01", assert)

		// Test just before the window
		monday1159am := time.Date(2016, 3, 21, 11, 59, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(monday1159am)

		assert.False(isOpen, "Should be closed before the window")
		assert.Equal(12, nextChange.Hour(), "Should open at 12:00")
		assert.Equal(0, nextChange.Minute(), "Should open at 12:00")
		assert.Equal(1*time.Minute, duration, "Should open in 1 minute")

		// Test during the window
		monday1200pm := time.Date(2016, 3, 21, 12, 0, 30, 0, time.UTC)
		isOpen, nextChange, duration = oh.MatchExt(monday1200pm)

		assert.True(isOpen, "Should be open during the window")
		assert.Equal(12, nextChange.Hour(), "Should close at 12:01")
		assert.Equal(1, nextChange.Minute(), "Should close at 12:01")
		assert.Equal(30*time.Second, duration, "Should close in 30 seconds")
	})

	t.Run("Month boundary with specific day", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("20 Mar Mo 10:00-16:00", assert)

		// March 20, 2016 was a Sunday, so this should be closed
		march20 := time.Date(2016, 3, 20, 12, 0, 0, 0, time.UTC)
		isOpen, nextChange, _ := oh.MatchExt(march20)

		assert.False(isOpen, "Should be closed on March 20 (Sunday)")
		// Next occurrence would be March 20, 2017 (if it's a Monday) or next available
		assert.False(nextChange.IsZero(), "Should have a next change time")
	})

	t.Run("Overlapping time ranges", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("Mo 09:00-15:00,12:00-18:00", assert)

		// Test during overlap period
		monday1pm := time.Date(2016, 3, 21, 13, 0, 0, 0, time.UTC)
		isOpen, nextChange, duration := oh.MatchExt(monday1pm)

		assert.True(isOpen, "Should be open during overlap")
		// Should close at the end of the longest range (18:00)
		assert.Equal(18, nextChange.Hour(), "Should close at 18:00")
		assert.Equal(5*time.Hour, duration, "Should close in 5 hours")
	})
}