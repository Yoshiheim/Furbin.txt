package handlers

import (
	"hoxt/data"
	"hoxt/internal/db"
	"hoxt/internal/helpers"
	"hoxt/internal/modules"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var tops []modules.Topic

	act := db.DB.Find(&tops)
	if act.Error != nil {
		log.Println(act.Error.Error())
		http.Error(w, "DB Error.", http.StatusInternalServerError)
		return
	}

	for i := range tops {
		tops[i].Name = html.EscapeString(tops[i].Name)
		tops[i].Description = html.EscapeString(tops[i].Description)
	}
	tpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, map[string]any{
		"data":   data.Configs,
		"logo":   html.EscapeString(string(data.Logo)),
		"topics": tops,
	}); err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}
}

/*
func FindByTopic(w http.ResponseWriter, r *http.Request) {

		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Query 'id' is empty", http.StatusBadRequest)
			return
		}
		log.Println(id)

		var pas []modules.Paste

		act := db.DB.Where("topic_id", id).Find(&pas)
		if act.Error != nil {
			http.Error(w, act.Error.Error(), http.StatusInternalServerError)
			return
		}

		for i := range pas {
			pas[i].Title = html.EscapeString(pas[i].Title)
			pas[i].Content = html.EscapeString(pas[i].Content)
		}

		tpl, err := template.ParseFiles("./templates/topicpastes.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tpl.Execute(w, map[string]any{
			"pastes": pas,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
*/
func FindByTopic(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "No topic id", http.StatusBadRequest)
		return
	}

	topicID, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid topic id", http.StatusBadRequest)
		return
	}

	var pas []modules.Paste

	act := db.DB.
		Preload("Topic").
		Where("topic_id = ?", topicID).
		Order("is_titled DESC").
		Find(&pas)

	if act.Error != nil {
		http.Error(w, act.Error.Error(), http.StatusInternalServerError)
		return
	}

	/*
		for i := range pas {
			pas[i].Title = html.EscapeString(pas[i].Title)
			pas[i].Content = html.EscapeString(pas[i].Content)
		}
	*/

	tpl, err := template.ParseFiles("./templates/topicpastes.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}

	temp, err := helpers.ParseCustomDuration(data.Configs.ClearTimer.Temp)
	if err != nil {
		temp = 0
	}

	tpl.Execute(w, map[string]any{
		"id":     topicID,
		"temp":   temp,
		"pastes": pas,
	})
}

func CreateTopic(w http.ResponseWriter, r *http.Request) {

}
