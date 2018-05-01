package main

import (
	"time"
)

type Birthday struct {
	date time.Time
	text string
}

func (b Birthday) NextBirthday() time.Time {
	return b.date
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
