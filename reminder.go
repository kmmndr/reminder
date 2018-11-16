package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/kmmndr/reminder/config"
	"github.com/kmmndr/reminder/reporter"
)

func postscriptum(reminderConf, binPath string) {
	fmt.Printf("ps.:\n")
	fmt.Printf(" - configuration %s\n", reminderConf)
	// fmt.Printf(" - binaire %s\n", "/home/bin/reminder.sh")
}

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func main() {
	var configFile = flag.String("reminder-conf", "reminder.conf", "Path to reminder.conf")
	flag.Parse()

	birthdays := config.ReadFile(*configFile)

	// now = birthdays[2].Time()
	now := time.Now().Local()
	now = Bod(now)
	// fmt.Printf("now %s\n", now)

	birthdays = birthdays.NextBirthdaysAfter(now)

	// fmt.Printf("%s\n", birthdays.String())

	// fmt.Printf("ref: %v\n", now)

	birthdays = birthdays.After(now)

	oneWeekFromNow := now.Add(time.Hour * 24 * 4)
	birthdays = birthdays.Before(oneWeekFromNow)

	// fmt.Printf("%s\n", birthdays.String())

	if birthdays.Len() > 0 {
		fmt.Printf("%s", reporter.Report(&birthdays, now))

		postscriptum(*configFile, "")
	} else {
		os.Exit(1)
	}
}
