package helpers

import (
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
}
