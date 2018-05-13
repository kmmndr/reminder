package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"kod.tapata.net/reminder/events"
)

func ReadFile(filename string) events.Birthdays {
	birthdays := make(events.Birthdays, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		r := regexp.MustCompile(`(?P<Year>\d{4}):(?P<Month>\d{2}):(?P<Day>\d{2}):(?P<Text>[^#]*)`)
		match := r.FindStringSubmatch(line)
		result := make(map[string]string)

		for i, name := range r.SubexpNames() {
			if i > 0 && i <= len(match) {
				result[name] = match[i]
			}
		}

		const shortForm = "2006-01-02"
		str := fmt.Sprintf("%4s-%2s-%2s", result["Year"], result["Month"], result["Day"])
		t, err := time.Parse(shortForm, str)
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
		birthdays = append(birthdays, events.NewBirthday(t, strings.TrimSpace(result["Text"])))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return birthdays
}
