package helpers

import (
	"html"
	"html/template"
	"strings"
	"time"
)

// Function For HTML-Templates on path '/HOXT/templates/*'
var FuncMap = template.FuncMap{
	"upper": strings.ToUpper,
	"formatDate": func(t time.Time) string {
		return t.Format("Monday, Jan 2, 2006 15:04:05")
	},
	"Escape": func(text string) string {
		return html.EscapeString(text)
	},
	"JoinEscape": func(text []string) string {
		var s []string
		for _, v := range text {
			s = append(s, EscapeString(v))
		}
		return strings.Join(s, "\n")
	},
}
