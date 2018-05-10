package reporter

import (
	"log"
	"strings"
	"time"

	"github.com/alecthomas/template"
	"kod.tapata.net/reminder/events"
)

const defaultTemplate = `###
## Notification des dates a venir prochainement
#
{{ range . }}
- {{ .Text }} ({{ format "02/01/2006" .Birthday }}) --> {{ comment . }}{{ end }}
`

func Report(birthdays *events.Birthdays, ref time.Time) string {
	var str strings.Builder

	funcs := template.FuncMap{
		"format": func(layout string, date time.Time) string {
			return date.Format(layout)
		},
		"comment": func(e events.Eventable) string {
			days := int(e.DaysBetween(ref))
			var str string

			switch days {
			case 0:
				str = "aujourd'hui !!!"
			case 1:
				str = "dans 1 jour ..."
			case 3:
				str = "dans 3 jours ..."
			}

			return str
		},
	}

	tpl := template.Must(template.New("birthdays").Funcs(funcs).Parse(defaultTemplate))

	err := tpl.Execute(&str, *birthdays)
	if err != nil {
		log.Fatal("executing template:", err)
	}

	return str.String()
}
