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
	"time"
)

type TopicWithCount struct {
	ID          uint
	Name        string
	Description string
	CreatedAt   time.Time
	PostCount   int
}

// The Index aka '/' path in website.
// path: 'http://<HOST>:<PORT>/'
func Index(w http.ResponseWriter, r *http.Request) {
	/*
		var tops []modules.Topic

		act := db.DB.Find(&tops)
		if act.Error != nil {
			log.Println(act.Error.Error())
			http.Error(w, "Error With DB.", http.StatusInternalServerError)
			return
		}
	*/

	var tops []TopicWithCount

	if err := db.DB.
		Model(&modules.Topic{}).
		Select(`
		topics.*,
		(SELECT COUNT(*) FROM pastes WHERE pastes.topic_id = topics.id) as post_count
	`).Scan(&tops); err.Error != nil {
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}

	for i := range tops {
		tops[i].Name = html.EscapeString(tops[i].Name)
		tops[i].Description = html.EscapeString(tops[i].Description)
	}
	tpl, err := template.New("index.html").Funcs(helpers.FuncMap).ParseFiles("./templates/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error With File", http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, map[string]any{
		"data":   data.Configs,
		"logo":   string(data.Logo),
		"topics": tops,
	}); err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}
}

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
		"topic":  topic,
	})
}
