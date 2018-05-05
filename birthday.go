package main

import (
	"time"
)

type Birthday struct {
	date time.Time
	text string
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

type Birthdays []Birthday

// Len is part of sort.Interface.
func (b Birthdays) Len() int {
	return len(b)
}

// Swap is part of sort.Interface.
func (b Birthdays) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less is part of sort.Interface.
func (b Birthdays) Less(i, j int) bool {
	return b[i].date.Before(b[j].date)
}
