package events

import (
	"time"
)

type BirthdayEvent struct {
	Event

	birthday time.Time
}

func NewBirthday(date time.Time, text string) BirthdayEvent {
	b := BirthdayEvent{
		birthday: date,
		Event: Event{
			text: text,
		},
	}
	return b
}

func (b *BirthdayEvent) BirthdayAfter(now time.Time) time.Time {
	year := now.Year()
	_, month, day := b.birthday.Date()
	dateCurrentYear := time.Date(year, month, day, 0, 0, 0, 0, b.birthday.Location())

	if now.After(dateCurrentYear) {
		return time.Date(year+1, month, day, 0, 0, 0, 0, b.birthday.Location())
	} else {
		return dateCurrentYear
	}
}

func (b *BirthdayEvent) NextBirthday() time.Time {
	now := time.Now()

	return b.BirthdayAfter(now)
}

func (b *BirthdayEvent) Next() BirthdayEvent {
	return BirthdayEvent{
		birthday: b.time,
		Event: Event{
			time: b.NextBirthday(),
			text: b.text,
		},
	}
}

func (b *BirthdayEvent) Birthdate() time.Time {
	return b.birthday
}
