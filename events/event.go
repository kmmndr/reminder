package events

import (
	"fmt"
	"time"
)

type Eventable interface {
	Text() string
	Time() time.Time
	Date() (int, time.Month, int)
	DaysBetween(ref time.Time) float64

	String() string
}

type Event struct {
	time time.Time
	text string
}

func New(time time.Time, text string) Event {
	e := Event{time, text}
	return e
}

func (e *Event) String() string {
	return fmt.Sprintf("%s: \"%s\"", e.time.String(), e.text)
}

func (e *Event) Time() time.Time {
	return e.time
}

func (e *Event) Date() (int, time.Month, int) {
	return e.Time().Date()
}

func (e *Event) Text() string {
	return e.text
}

func (e *Event) TimeEqual(ref time.Time) bool {
	utc, _ := time.LoadLocation("UTC")

	return e.time.In(utc).Equal(ref.In(utc))
}

func (e *Event) DaysBetween(ref time.Time) float64 {
	diff := -ref.Sub(e.Time())
	days := diff.Hours() / 24

	return days
}

func (e *Event) DaysFromNow() float64 {
	return e.DaysBetween(time.Now())
}
