package handlers

import (
	"hoxt/internal/db"
	"hoxt/internal/modules"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

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

	tpl, err := template.ParseFiles("./templates/local.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}

	var paste modules.Paste

	act := db.DB.Find(&paste, id)
	if act.Error != nil {
		log.Println(err.Error())
		http.Error(w, act.Error.Error(), http.StatusInternalServerError)
		return
	}

	//i would prefer to escape this.
	paste.Title = html.EscapeString(paste.Title)
	paste.Content = html.EscapeString(paste.Content)
	paste.Author = html.EscapeString(paste.Author)

	if err := tpl.Execute(w, paste); err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}
}
