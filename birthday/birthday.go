package birthday

import (
	"fmt"
	"time"
)

type Birthday struct {
	date time.Time
	text string
}

func New(date time.Time, text string) Birthday {
	b := Birthday{date, text}
	return b
}

func (b Birthday) BirthdayAfter(now time.Time) time.Time {
	year := now.Year()
	_, month, day := b.date.Date()
	dateCurrentYear := time.Date(year, month, day, 0, 0, 0, 0, b.date.Location())

	if now.After(dateCurrentYear) {
		return time.Date(year+1, month, day, 0, 0, 0, 0, b.date.Location())
	} else {
		return dateCurrentYear
	}
}

func (b Birthday) NextBirthday() time.Time {
	now := time.Now()

	return b.BirthdayAfter(now)
}

func (b Birthday) Next() Birthday {
	return Birthday{date: b.NextBirthday(), text: b.text}
}

func (b Birthday) String() string {
	return fmt.Sprintf("%s: \"%s\"", b.date.String(), b.text)
}

func (b Birthday) Time() time.Time {
	return b.date
}

func (b Birthday) Date() (int, time.Month, int) {
	return b.date.Date()
}

func (b Birthday) Text() string {
	return b.text
}

func (b Birthday) DateIs(ref time.Time) bool {
	utc, _ := time.LoadLocation("UTC")

	return b.date.In(utc).Equal(ref.In(utc))
}
