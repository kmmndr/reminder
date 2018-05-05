package birthday

import (
	"sort"
	"strings"
	"time"
)

type Birthdays []Birthday

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
	return bs[i].date.Before(bs[j].date)
}

func (bs *Birthdays) NextBirthdays() Birthdays {
	nextBirthdays := make(Birthdays, len(*bs))

	for idx, birthday := range *bs {
		nextBirthdays[idx] = birthday.Next()
	}

	sort.Sort(nextBirthdays)

	return nextBirthdays
}

func (bs *Birthdays) Before(ref time.Time) Birthdays {
	birthdays := make(Birthdays, 0)

	for _, birthday := range *bs {
		if birthday.DateIs(ref) || birthday.date.Before(ref) {
			birthdays = append(birthdays, birthday)
		}
	}

	return birthdays
}

func (bs *Birthdays) After(ref time.Time) Birthdays {
	birthdays := make(Birthdays, 0)

	for _, birthday := range *bs {
		if birthday.DateIs(ref) || birthday.date.After(ref) {
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
