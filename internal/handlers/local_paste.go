package handlers

import (
	"hoxt/internal/db"
	"hoxt/internal/helpers"
	"hoxt/internal/modules"
	"html"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Some Local Paste Website.
// Path: 'http://<HOST>:<PORT>/paste/1'
func Local(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "No topic id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid topic id", http.StatusBadRequest)
		return
	}

	var paste modules.Paste
	var count int64

	act := db.DB.Find(&paste, id).Count(&count)
	if act.Error != nil {
		log.Println(err.Error())
		http.Error(w, act.Error.Error(), http.StatusInternalServerError)
		return
	}

	if count <= 0 {
		helpers.Render404(w)
		return
	}

	//i would prefer to escape this.
	paste.Title = html.EscapeString(paste.Title)
	paste.Content = html.EscapeString(paste.Content)
	paste.Author = html.EscapeString(paste.Author)

	tpl, err := template.New("local.html").Funcs(helpers.FuncMap).ParseFiles("./templates/local.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, paste); err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}
}
