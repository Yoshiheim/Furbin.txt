package handlers

import (
	"hoxt/data"
	"hoxt/internal/db"
	"hoxt/internal/helpers"
	"hoxt/internal/modules"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Index Of Topic of Website.
// Path: 'http://<HOST>:<PORT>/topic/1'
func FindByTopic(w http.ResponseWriter, r *http.Request) {
	// Split "/topic/<int>" by '/' char
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "No topic id", http.StatusBadRequest)
		return
	}

	// parse <int>
	topicID, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid topic id", http.StatusBadRequest)
		return
	}

	// prepare vars for output
	var pas []modules.Paste
	var topic modules.Topic

	// sqlite query
	act := db.DB.
		Preload("Topic").
		Where("topic_id = ?", topicID).
		Order("is_titled DESC").
		Find(&pas)

	// catch error if its exist
	if act.Error != nil {
		http.Error(w, "something does wrong...", http.StatusInternalServerError)
		return
	}

	// second sqlite query
	act2 := db.DB.
		Where("id = ?", topicID).
		Find(&topic)

	// catch error if its exist for act2
	if act2.Error != nil {
		http.Error(w, "something does wrong...", http.StatusInternalServerError)
		return
	}

	// check is topic empty(because its doesn't exist)
	if topic == (modules.Topic{}) {
		// Render "/HOXT/templates/404.html" template
		helpers.Render404(w)
		return
	}

	// Escape all pastes
	// Its already escaped so why its here?
	/*
		for i := range pas {
			pas[i].Title = html.EscapeString(pas[i].Title)
			pas[i].Content = html.EscapeString(pas[i].Content)
			pas[i].Author = html.EscapeString(pas[i].Author)
		}
	*/

	// parse ClearTimer
	temp, err := helpers.ParseCustomDuration(data.Configs.ClearTimer.Temp)
	if err != nil {
		temp = 0
	}

	// parse "/HOXT/templates/topicpastes.html"

	tpl, err := template.New("topicpastes.html").Funcs(helpers.FuncMap).ParseFiles("./templates/topicpastes.html", "./templates/attr.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}

	//and render for client
	tpl.Execute(w, map[string]any{
		"id":     topicID,
		"temp":   temp,
		"pastes": pas,
		"topic":  topic,
	})
}
