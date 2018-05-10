package events

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const shortForm = "2006-01-02"

func TestBirthdayEvent(t *testing.T) {
	ref, _ := time.Parse(shortForm, "2017-06-01")

	var dateTest = []struct {
		birthday     string
		nextBirthday string
		daysBetween  int
	}{
		{"2017-05-01", "2018-05-01", 334},
		{"2017-07-01", "2017-07-01", 30},
	}

	for _, tt := range dateTest {
		date, _ := time.Parse(shortForm, tt.birthday)
		birthdayEvent := NewBirthday(date, fmt.Sprintf("BirthdayEvent %s", tt.birthday))

		// year, month, day := birthday.BirthdayAfter(ref).Date()
		next := birthdayEvent.NextAfter(ref)

		year, month, day := next.Time().Date()
		nextBirthday := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
		assert.Equal(t, tt.nextBirthday, nextBirthday)

		year, month, day = next.Birthday().Date()
		birthday := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
		assert.Equal(t, tt.birthday, birthday)

		days := next.DaysBetween(ref)
		assert.Equal(t, tt.daysBetween, int(days))
	}
}
