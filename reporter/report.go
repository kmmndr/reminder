package reporter

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/alecthomas/template"
	"github.com/kmmndr/reminder/events"
)

const defaultTemplate = `###
## Notification des dates a venir prochainement ;-)
#

{{ range . }}{{ birthday . }}{{ end }}
`

func Report(birthdays *events.Birthdays, ref time.Time) string {
	var str strings.Builder

	funcs := template.FuncMap{
		"birthday": func(b *events.BirthdayEvent) string {
			var builder strings.Builder
			days := b.DaysBetween(ref)
			comment := ""

			switch int(days) {
			case 0:
				comment = "aujourd'hui !!!"
			case 1:
				comment = "dans 1 jour ..."
			default:
				comment = fmt.Sprintf("dans %d jours ...", int(days))
			}

			fmt.Fprintf(&builder, "%s (%s) --> %s\n", b.Text(), b.Birthday().Format("02/01/2006"), comment)

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
