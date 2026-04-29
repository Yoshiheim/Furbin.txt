package handlers

import (
	"hoxt/data"
	"hoxt/internal/db"
	"hoxt/internal/helpers"
	"hoxt/internal/modules"
	"html"
	"log"
	"net/http"
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
func MainPage(w http.ResponseWriter, r *http.Request) {
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

	tpl, err := template.New("index.html").Funcs(helpers.FuncMap).ParseFiles("./templates/index.html", "./templates/attr.html")
	if err != nil {
		http.Error(w, "Error With File", http.StatusInternalServerError)
		return
	}

	helpers.UpdateLogo()

	cfg := data.GetDConfig(w)

	if err := tpl.Execute(w, map[string]any{
		"data":   cfg,
		"logo":   string(data.Logo),
		"topics": tops,
	}); err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}
}
