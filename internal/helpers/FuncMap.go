package helpers

import (
	"fmt"
	"hoxt/data"
	"html"
	"html/template"
	"math/rand"
	"strconv"
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
	"CutString": func(text string, max int) string {
		return TruncateByte(text, max)
	},
	"JoinEscape": func(text []string) string {
		var s []string
		for _, v := range text {
			s = append(s, html.EscapeString(v))
		}
		return strings.Join(s, "\n")
	},
	"Sub": func(a, b int) int {
		return a - b
	},
	"Rand": func(a int) int {
		return rand.Intn(a)
	},
	"RGB2String": func(r data.RGB) string {
		return fmt.Sprintf("%s,%s,%s", strconv.Itoa(r.R), strconv.Itoa(r.G), strconv.Itoa(r.B))
	},
}
