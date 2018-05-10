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
## Notification des dates a venir prochainement
#
{{ range . }}{{ birthday . }}{{ end }}
`

func Report(birthdays *events.Birthdays, ref time.Time) string {
	var str strings.Builder

	funcs := template.FuncMap{
		"birthday": func(b *events.BirthdayEvent) string {
			var builder strings.Builder
			days := int(b.DaysBetween(ref))

			comments := make(map[int]string)
			comments[0] = "aujourd'hui !!!"
			comments[1] = "dans 1 jour ..."
			comments[3] = "dans 3 jours ..."

			if days == 0 || days == 1 || days == 3 {
				fmt.Fprintf(&builder, "%s (%s) --> %s\n", b.Text(), b.Birthday().Format("02/01/2006"), comments[days])
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
