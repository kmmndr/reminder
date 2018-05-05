package main

import (
	"flag"
	"fmt"
	"time"

	"kod.tapata.net/reminder/config"
	"kod.tapata.net/reminder/reporter"
)

func main() {
	var configFile = flag.String("config", "reminder.conf", "Path to reminder.conf")
	// var before = flag.Int("before", 7, "Before")
	flag.Parse()

	birthdays := config.ReadFile(*configFile)
	birthdays = birthdays.NextBirthdays()

	fmt.Printf("%s\n", birthdays.String())

	//now := time.Now()
	now := birthdays[2].Time()
	fmt.Printf("ref: %v\n", now)

	birthdays = birthdays.After(now)

	oneWeekFromNow := now.Add(time.Hour * 24 * 7)
	birthdays = birthdays.Before(oneWeekFromNow)

	fmt.Printf("%s\n", birthdays.String())

	fmt.Printf("Report: %s\n", reporter.Report(&birthdays))
}
