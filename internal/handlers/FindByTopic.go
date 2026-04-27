package handlers

import (
	"hoxt/data"
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

// Index Of Topic of Website.
// Path: 'http://<HOST>:<PORT>/topic/1'
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
	var topic modules.Topic

	act := db.DB.
		Preload("Topic").
		Where("topic_id = ?", topicID).
		Order("is_titled DESC").
		Find(&pas)

	if act.Error != nil {
		http.Error(w, "something does wrong...", http.StatusInternalServerError)
		return
	}

	act2 := db.DB.
		Where("id = ?", topicID).
		Find(&topic)

	if act2.Error != nil {
		http.Error(w, "something does wrong...", http.StatusInternalServerError)
		return
	}

	if topic == (modules.Topic{}) {
		helpers.Render404(w)
		return
	}

	for i := range pas {
		pas[i].Title = html.EscapeString(pas[i].Title)
		pas[i].Content = html.EscapeString(pas[i].Content)
		pas[i].Author = html.EscapeString(pas[i].Author)
	}

	tpl, err := template.New("topicpastes.html").Funcs(helpers.FuncMap).ParseFiles("./templates/topicpastes.html")
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
		"topic":  topic,
	})
}
