package openhours

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"math/rand"
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

func compileSuccess(s string, a *assert.Assertions) *OpenHours {
	oh, errs := CompileOpenHours(s)
	a.Empty(errs)
	fmt.Println(oh)

	return oh
}

func TestMatch(t *testing.T) {
	assert := assert.New(t)

	//weekdays only
	for wd, i := range weekdays {
		t := sunday().AddDate(0, 0, i)
		fmt.Println(t)

		oh := compileSuccess(wd, assert)

		assert.True(oh.Match(t))
		assert.False(oh.Match(t.Add(minusMinute)))

		assert.False(oh.Match(t.AddDate(0, 0, 1)))
		assert.True(oh.Match(t.AddDate(0, 0, 1).Add(minusMinute)))
	}

	//day with month
	{
		oh := compileSuccess("20 Mar", assert)
		march20 := date("20-03-2016")
		march21 := date("21-03-2016")

		assert.True(oh.Match(march20))
		assert.False(oh.Match(march20.Add(minusMinute)))

		assert.False(oh.Match(march21))
		assert.True(oh.Match(march21.Add(minusMinute)))

		assert.True(oh.Match(march20.AddDate(rand.Intn(20), 0, 0)))
		assert.False(oh.Match(march20.AddDate(0, rand.Intn(20), 0)))
	}

	//day only
	{
		oh := compileSuccess("20", assert)

		assert.False(oh.Match(date("19-01-2016")))
		assert.False(oh.Match(date("20-01-2016").Add(minusMinute)))
		assert.True(oh.Match(date("20-01-2016")))
		assert.True(oh.Match(date("21-01-2016").Add(minusMinute)))
		assert.False(oh.Match(date("21-01-2016")))

		for i := 1; i < 12; i++ {
			assert.True(oh.Match(sunday().AddDate(i, i, 0))) //any month or year
		}
	}

	//days range
	{
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
	}

	{
		oh := compileSuccess("20,21 Mar Mo", assert)

		assert.False(oh.Match(date("20-03-2016")))
		assert.True(oh.Match(date("21-03-2016")))
	}

	{
		oh := compileSuccess("06:30-10:00,17:00-18:30", assert)
		midnight := sunday()

		assert.False(oh.Match(midnight))
		assert.False(oh.Match(midnight.Add(6 * time.Hour)))

		assert.True(oh.Match(midnight.Add(6*time.Hour + 30*time.Minute)))

		t := time.Now()
		count := 0
		for i := 0; i < 1440; i++ {
			t = t.Add(time.Minute)
			if oh.Match(t) {
				count++
			}
		}
		assert.Equal(300, count)
	}
}

func TestCompilationFailure(t *testing.T) {
	assert := assert.New(t)

	_, errs := CompileOpenHours("")
	assert.NotEmpty(errs)

	_, errs = CompileOpenHours("-1")
	assert.NotEmpty(errs)

	_, errs = CompileOpenHours("0")
	assert.NotEmpty(errs)

	_, errs = CompileOpenHours("32")
	assert.NotEmpty(errs)

	_, errs = CompileOpenHours("20 March")
	assert.NotEmpty(errs)

	_, errs = CompileOpenHours("March")
	assert.NotEmpty(errs)

	_, errs = CompileOpenHours("20-03")
	assert.NotEmpty(errs)

	_, errs = CompileOpenHours("20/03")
	assert.NotEmpty(errs)

	_, errs = CompileOpenHours("20 Mar Xx")
	assert.NotEmpty(errs)

	_, errs = CompileOpenHours("20 Mar 10:00")
	assert.NotEmpty(errs)

	_, errs = CompileOpenHours("00:00-25:00")
	assert.NotEmpty(errs)

	_, errs = CompileOpenHours("We-Tu")
	assert.NotEmpty(errs)

	_, errs = CompileOpenHours("00:00:00-00:00:59")
	assert.NotEmpty(errs)

	_, errs = CompileOpenHours("14:00-11:00")
	assert.NotEmpty(errs)
}
