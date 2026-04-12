package handlers

import (
	"encoding/json"
	"hoxt/internal/db"
	"hoxt/internal/helpers"
	"hoxt/internal/modules"
	"html"
	"log"
	"net/http"
	"strings"
	"text/template"
)

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

func CreatePaste(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Author  string `json:"author"`
		TopicID uint   `json:"topicid"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println("JSON Decode Error:", err)
		log.Printf("%s - %s - %d", body.Title, body.Content, body.TopicID)
		http.Error(w, "Cannot parse JSON body", http.StatusBadRequest)
		return
	}
	if body.Title == "" || body.Content == "" {
		log.Println("body is empty")
		log.Printf("%s - %s - %d", body.Title, body.Content, body.TopicID)
		http.Error(w, "body is empty", http.StatusBadRequest)
		return
	}
	body.Title = html.EscapeString(helpers.SanitizeString(helpers.TruncateByte(body.Title, 100)))
	body.Content = html.EscapeString(helpers.SanitizeString(helpers.TruncateByte(body.Content, 50000)))
	body.Author = html.EscapeString(helpers.SanitizeString(helpers.TruncateByte(body.Author, 100)))

	act := db.DB.Create(&modules.Paste{
		Title:   body.Title,
		Content: body.Content,
		Author:  body.Author,
		TopicID: body.TopicID,
	})
	if act.Error != nil {
		if strings.Contains(act.Error.Error(), "violates foreign key constraint") {
			log.Println(body.TopicID)
			http.Error(w, "Указанный родительский ID не существует", http.StatusBadRequest)
			return
		}
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
}
