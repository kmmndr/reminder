package reporter

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/alecthomas/template"
	"kod.tapata.net/reminder/birthday"
)

const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

const defaultTemplate = `
###
## Notification des dates a venir prochainement
#

{{ range . }}
- {{ .Text }} ({{ format "02/01" .Time }})
{{ end }}

// Numero 3 ## (01/05/2016) --> aujourd'hui !!!
// Numero 1 ## (02/05/2016) --> dans 1 jour(s) ...

ps.:
 - configuration /home/samba/nous/reminder.conf
 - binaire /home/bin/reminder.sh
`

/*

 */

func Report(birthdays *birthday.Birthdays) string {
	var str strings.Builder

	funcs := template.FuncMap{"format": func(layout string, date time.Time) string { return date.Format(layout) }}
	tpl := template.Must(template.New("birthdays").Funcs(funcs).Parse(defaultTemplate))

	err := tpl.Execute(os.Stdout, *birthdays)
	if err != nil {
		log.Fatal("executing template:", err)
	}

	//str.WriteString(birthday.String())
	//str.WriteString("\n")
	//}

	return str.String()
}
