package openhours

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var minusMinute = -1 * time.Minute

func date(s string) time.Time {
	t, err := time.Parse("02-01-2006", s)
	if err != nil {
		panic(err)
	}
	return t
}

func sunday() time.Time {
	return time.Date(2016, time.March, 20, 0, 0, 0, 0, time.UTC)
}

func compileSuccess(s string, a *require.Assertions) *OpenHours {
	oh, errs := CompileOpenHours(s)
	a.Empty(errs, fmt.Sprintf("errors %v on compiling %s", errs, s))
	a.NotNil(oh, fmt.Sprintf("nil result on compiling %s", s))

	return oh
}

func TestMatch(t *testing.T) {
	t.Run("24/7", func(t *testing.T) {
		assert := require.New(t)
		oh := compileSuccess("24/7", assert)

		assert.True(oh.Match(time.Now()))
		assert.True(oh.Match(date("01-01-1970")))
		assert.True(oh.Match(date("22-11-2028")))
	})

	t.Run("weekdays only", func(t *testing.T) {
		weekdays := map[string]int{"Su": 0, "Mo": 1, "Tu": 2, "We": 3, "Th": 4, "Fr": 5, "Sa": 6}
		for wd, i := range weekdays {
			t.Run(wd, func(t *testing.T) {
				assert := require.New(t)

				date := sunday().AddDate(0, 0, i)
				fmt.Println(date)

				oh := compileSuccess(wd, assert)

				assert.True(oh.Match(date))
				assert.False(oh.Match(date.Add(minusMinute)))

				assert.False(oh.Match(date.AddDate(0, 0, 1)))
				assert.True(oh.Match(date.AddDate(0, 0, 1).Add(minusMinute)))
			})
		}
	})

	t.Run("day with month", func(t *testing.T) {
		assert := require.New(t)

		oh := compileSuccess("20 Mar", assert)
		march20 := date("20-03-2016")
		march21 := date("21-03-2016")

		assert.True(oh.Match(march20))
		assert.False(oh.Match(march20.Add(minusMinute)))

		assert.False(oh.Match(march21))
		assert.True(oh.Match(march21.Add(minusMinute)))

		assert.True(oh.Match(march20.AddDate(5, 0, 0)))
		assert.False(oh.Match(march20.AddDate(0, 3, 0)))
	})

	t.Run("day only", func(t *testing.T) {
		assert := require.New(t)

		oh := compileSuccess("20", assert)

		assert.False(oh.Match(date("19-01-2016")))
		assert.False(oh.Match(date("20-01-2016").Add(minusMinute)))
		assert.True(oh.Match(date("20-01-2016")))
		assert.True(oh.Match(date("21-01-2016").Add(minusMinute)))
		assert.False(oh.Match(date("21-01-2016")))

		for i := 1; i < 12; i++ {
			assert.True(oh.Match(sunday().AddDate(i, i, 0))) //any month or year
		}
	})

	t.Run("day range", func(t *testing.T) {
		assert := require.New(t)

		oh := compileSuccess("20-23", assert)

		assert.False(oh.Match(date("19-01-2016")))
		assert.False(oh.Match(date("20-01-2016").Add(minusMinute)))
		assert.True(oh.Match(date("20-01-2016")))
		assert.True(oh.Match(date("22-01-2016")))
		assert.True(oh.Match(date("23-01-2016")))
		assert.True(oh.Match(date("24-01-2016").Add(minusMinute)))
		assert.False(oh.Match(date("24-01-2016")))

		for i := 1; i < 3; i++ {
			t := sunday().AddDate(i, i, i)
			assert.True(oh.Match(t))
			assert.True(oh.Match(t.Add(minusMinute)))
		}
	})

	t.Run("day + month + weekday", func(t *testing.T) {
		assert := require.New(t)

		oh := compileSuccess("20,21 Mar Mo", assert)

		assert.False(oh.Match(date("20-03-2016")))
		assert.True(oh.Match(date("21-03-2016")))
	})

	t.Run("time intervals only", func(t *testing.T) {
		assert := require.New(t)

		oh := compileSuccess("06:30-10:00,17:00-18:30", assert)
		midnight := sunday()

		assert.False(oh.Match(midnight))
		assert.False(oh.Match(midnight.Add(6 * time.Hour)))

		assert.True(oh.Match(midnight.Add(6*time.Hour + 30*time.Minute)))

		ti := time.Now()
		count := 0
		for i := 0; i < 1440; i++ {
			ti = ti.Add(time.Minute)
			if oh.Match(ti) {
				count++
			}
		}
		assert.Equal(300, count)
	})

	t.Run("semicolon", func(t *testing.T) {
		type testCase struct {
			name string
			def  string
		}

		for _, tc := range []testCase{
			{"no sep", "Mo 10:00-12:00;Tu 10:00-12:00"},
			{"single space", "Mo 10:00-12:00; Tu 10:00-12:00"},
			{"multiple space", "Mo 10:00-12:00;    Tu 10:00-12:00"},
			{"signele tab", "Mo 10:00-12:00;	Tu 10:00-12:00"},
			{"multiple tab", "Mo 10:00-12:00;				Tu 10:00-12:00"},
		} {
			t.Run(tc.name, func(t *testing.T) {
				assert := require.New(t)
				oh := compileSuccess(tc.def, assert)

				assert.False(oh.Match(date("03-12-2017").Add(10 * time.Hour)))
				assert.True(oh.Match(date("04-12-2017").Add(10 * time.Hour)))
			})
		}
	})
}

func TestCompilationFailure(t *testing.T) {
	for i, tc := range []string{"", "-1", "0", "32", "20 March", "March", "20-03", "20/03", "20 Mar Xx", "20 Mar 10:00", "00:00-25:00", "We-Tu", "00:00:00-00:00:59", "14:00-11:00"} {
		t.Run(fmt.Sprintf("%d.%s", i, tc), func(t *testing.T) {
			assert := assert.New(t)

			_, errs := CompileOpenHours("")
			assert.NotEmpty(errs)
		})
	}
}
