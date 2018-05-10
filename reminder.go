package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"kod.tapata.net/reminder/config"
	"kod.tapata.net/reminder/reporter"
)

func postscriptum(reminderConf, binPath string) {
	fmt.Printf("ps.:\n")
	fmt.Printf(" - configuration %s\n", reminderConf)
	// fmt.Printf(" - binaire %s\n", "/home/bin/reminder.sh")
}

func main() {
	var configFile = flag.String("reminder-conf", "reminder.conf", "Path to reminder.conf")
	flag.Parse()

	birthdays := config.ReadFile(*configFile)
	birthdays = birthdays.NextBirthdays()

	// fmt.Printf("%s\n", birthdays.String())

	now := time.Now()
	// now = birthdays[2].Time()
	// fmt.Printf("ref: %v\n", now)

	birthdays = birthdays.After(now)

	oneWeekFromNow := now.Add(time.Hour * 24 * 3)
	birthdays = birthdays.Before(oneWeekFromNow)

	// fmt.Printf("%s\n", birthdays.String())

	if birthdays.Len() > 0 {
		fmt.Println(reporter.Report(&birthdays, now))

		postscriptum(*configFile, "")
	} else {
		os.Exit(1)
	}
}
