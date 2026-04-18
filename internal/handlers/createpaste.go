package handlers

import (
	"encoding/json"
	"hoxt/internal/db"
	"hoxt/internal/helpers"
	"hoxt/internal/modules"
	"html"
	"net/http"
	"strings"
)

/*
func CreatePasteHtmlTemplate(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./templates/create_paste.html")
	if err != nil {
		http.Error(w, "Cant Get Template", http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, nil); err != nil {
		http.Error(w, "Cant Parse Template", http.StatusInternalServerError)
		return
	}
}
*/

func CreatePaste(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Author  string `json:"author"`
		TopicID uint   `json:"topicid"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Cannot parse JSON body", http.StatusBadRequest)
		return
	}

	if len(body.Title) > 128 {
		http.Error(w, "Title text-field exceeds character limit of 128.", http.StatusBadRequest)
		return
	}

	// 64kb text limit
	if len(body.Content) > 65536 {
		http.Error(w, "Content text-field exceeds character limit of 65536.", http.StatusBadRequest)
		return
	}

	if len(body.Author) > 128 {
		http.Error(w, "Author text-field exceeds character limit of 128.", http.StatusBadRequest)
		return
	}

	body.Title = html.EscapeString(helpers.SanitizeString(body.Title))
	body.Content = html.EscapeString(helpers.SanitizeString(body.Content))
	body.Author = html.EscapeString(helpers.SanitizeString(strings.ReplaceAll(body.Author, " ", "")))

	if body.Title == "" {
		http.Error(w, "Title Is empty", http.StatusBadRequest)
		return
	}

	if body.Content == "" {
		http.Error(w, "Content Is empty", http.StatusBadRequest)
		return
	}

	act := db.DB.Create(&modules.Paste{
		Title:   body.Title,
		Content: body.Content,
		Author:  body.Author,
		TopicID: body.TopicID,
	})

	if act.Error != nil {
		if strings.Contains(act.Error.Error(), "violates foreign key constraint") {
			http.Error(w, "Topic ID does not exist ", http.StatusBadRequest)
			return
		}

		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
}
