package reporter

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/alecthomas/template"
	"kod.tapata.net/reminder/events"
)

const defaultTemplate = `###
## Notification des dates a venir prochainement ;-)
#

{{ range . }}{{ birthday . }}{{ end }}
`

func daysBetween(value, min float64) bool {
	max := float64(int(min) + 1)
	return min <= value && value < max
}

func Report(birthdays *events.Birthdays, ref time.Time) string {
	var str strings.Builder

	funcs := template.FuncMap{
		"birthday": func(b *events.BirthdayEvent) string {
			var builder strings.Builder
			days := b.DaysBetween(ref)

			comments := make(map[int]string)
			comments[0] = "aujourd'hui !!!"
			comments[1] = "dans 1 jour ..."
			comments[3] = "dans 3 jours ..."

			if daysBetween(days, 0) || daysBetween(days, 1) || daysBetween(days, 3) {
				fmt.Fprintf(&builder, "%s (%s) --> %s\n", b.Text(), b.Birthday().Format("02/01/2006"), comments[int(days)])
			}

			return builder.String()
		},
	}

	tpl := template.Must(template.New("birthdays").Funcs(funcs).Parse(defaultTemplate))

	err := tpl.Execute(&str, *birthdays)
	if err != nil {
		log.Fatal("executing template:", err)
	}

	return str.String()
}
