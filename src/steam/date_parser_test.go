package steam_test

import (
	"github.com/stretchr/testify/assert"
	"go.llib.dev/testcase"
	"steam-screenshot-cli/src/steam"
	"testing"
	"time"
)

var testingMode bool

var dateStringWithYear = "8 Aug, 2023 @ 9:57pm"
var dateWithYear = time.Date(2023, 8, 8, 21, 57, 0, 0, time.UTC)

var dateStringWithoutYear = "6 Jan @ 4:07am"

var dateWithoutYear = time.Date(time.Now().Year(), 1, 6, 4, 7, 0, 0, time.UTC)

var badDate1 = "kekew"
var badDate2 = "kekew,"

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestDates(t *testing.T) {

	date, err := steam.ParseSteamDate(dateStringWithYear)

	assert.Equal(t, nil, err, "No Error")
	assert.Equal(t, dateWithYear, date, "Dates With Year Works")

	date, err = steam.ParseSteamDate(dateStringWithoutYear)
	assert.Equal(t, nil, err, "No Error")
	assert.Equal(t, dateWithoutYear, date, "Dates With Year Works")

	date, err = steam.ParseSteamDate(badDate1)
	assert.NotEqual(t, nil, err, "Error")
	assert.NotEqual(t, dateWithoutYear, date, "Dates not works")

}

func TestSpecParseDate(t *testing.T) {
	s := testcase.NewSpec(t)
	s.NoSideEffect()

	var (
		input = testcase.Let[string](s, nil)
	)
	act := func(t *testcase.T) (time.Time, error) {
		return steam.ParseSteamDate(input.Get(t))
	}

	s.When("Date Value with year", func(s *testcase.Spec) {
		// testing scopes without the input being defined
		// will warn the dev about the missing specification.
		input.LetValue(s, dateStringWithYear)

		s.Then("it is expected to work", func(t *testcase.T) {
			date, err := act(t)
			t.Must.Equal(nil, err)
			t.Must.Equal(dateWithYear, date)
		})
	})

	s.When("Date Value without year year", func(s *testcase.Spec) {
		input.LetValue(s, dateStringWithoutYear)

		s.Then("it is expected to work", func(t *testcase.T) {
			date, err := act(t)
			t.Must.Equal(nil, err)
			t.Must.Equal(dateWithoutYear, date)
		})
	})

	s.When("Date Value is Bad date", func(s *testcase.Spec) {
		input.LetValue(s, badDate1)
		input.LetValue(s, badDate2)
		s.Then("it is expected not to work 1", func(t *testcase.T) {

			date, err := act(t)
			t.Must.Error(err)
			t.Must.Equal(time.Time{}, date)
		})

	})
}
