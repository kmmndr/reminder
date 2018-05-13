package events

import (
	"sort"
	"strings"
	"time"
)

type Birthdays []BirthdayEvent

// Len is part of sort.Interface.
func (bs Birthdays) Len() int {
	return len(bs)
}

// Swap is part of sort.Interface.
func (bs Birthdays) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

// Less is part of sort.Interface.
func (bs Birthdays) Less(i, j int) bool {
	return bs[i].time.Before(bs[j].time)
}

func (bs *Birthdays) NextBirthdaysAfter(ref time.Time) Birthdays {
	nextBirthdays := make(Birthdays, len(*bs))

	for idx, birthday := range *bs {
		nextBirthdays[idx] = birthday.NextAfter(ref)
	}

	sort.Sort(nextBirthdays)

	return nextBirthdays
}

func (bs *Birthdays) NextBirthdays() Birthdays {
	return bs.NextBirthdaysAfter(time.Now())
}

func (bs *Birthdays) Before(ref time.Time) Birthdays {
	birthdays := make(Birthdays, 0)

	for _, birthday := range *bs {
		if birthday.TimeEqual(ref) || birthday.time.Before(ref) {
			birthdays = append(birthdays, birthday)
		}
	}

	return birthdays
}

func (bs *Birthdays) After(ref time.Time) Birthdays {
	birthdays := make(Birthdays, 0)

	for _, birthday := range *bs {
		if birthday.TimeEqual(ref) || birthday.time.After(ref) {
			birthdays = append(birthdays, birthday)
		}
	}

	return birthdays
}

func (bs *Birthdays) String() string {
	var str strings.Builder

	for _, birthday := range *bs {
		str.WriteString(birthday.String())
		str.WriteString("\n")
	}

	return str.String()
}
