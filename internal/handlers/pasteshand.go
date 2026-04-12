package handlers

import (
	"hoxt/data"
	"hoxt/internal/db"
	"hoxt/internal/modules"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func IndexPastes(w http.ResponseWriter, r *http.Request) {

	var pas []modules.Paste

	act := db.DB.Order("is_titled DESC").Find(&pas)
	if act.Error != nil {
		http.Error(w, act.Error.Error(), http.StatusInternalServerError)
		return
	}

	for i := range pas {
		pas[i].Title = html.EscapeString(pas[i].Title)
		pas[i].Content = html.EscapeString(pas[i].Content)
	}

	tpl, err := template.ParseFiles("./templates/pastes.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, map[string]any{
		"data":   data.Configs,
		"logo":   data.Logo,
		"pastes": pas,
	}); err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}
}

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
		http.Error(w, act.Error.Error(), http.StatusInternalServerError)
		return
	}

	//paste.Title = html.EscapeString(paste.Title)
	//paste.Content = html.EscapeString(paste.Content)

	if err := tpl.Execute(w, paste); err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}
}
