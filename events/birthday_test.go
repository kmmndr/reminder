package events

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const shortForm = "2006-01-02"

func TestBirthdayEvent(t *testing.T) {
	var dateTest = []struct {
		birthday     string
		nextBirthday string
	}{
		{"2017-05-01", "2018-05-01"},
		{"2017-07-01", "2017-07-01"},
	}

	ref, _ := time.Parse(shortForm, "2017-06-01")
	for _, tt := range dateTest {
		date, _ := time.Parse(shortForm, tt.birthday)
		birthday := NewBirthday(date, fmt.Sprintf("BirthdayEvent %s", tt.birthday))

		year, month, day := birthday.BirthdayAfter(ref).Date()
		nextBirthday := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
		assert.Equal(t, tt.nextBirthday, nextBirthday)
	}
}
