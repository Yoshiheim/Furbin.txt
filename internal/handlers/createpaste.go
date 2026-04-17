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
		log.Printf("JSON Error Decode: %q - %q - %q: %q", body.Title, body.Content, body.TopicID, err.Error())
		http.Error(w, "Cannot parse JSON body", http.StatusBadRequest)
		return
	}

	body.Title = html.EscapeString(helpers.SanitizeString(helpers.TruncateByte(body.Title, 100)))
	body.Content = html.EscapeString(helpers.SanitizeString(helpers.TruncateByte(body.Content, 10000)))
	body.Author = html.EscapeString(helpers.SanitizeString(helpers.TruncateByte(strings.ReplaceAll(body.Author, " ", ""), 100)))

	if body.Title == "" || body.Content == "" {
		log.Printf("Some Body is Empty: %q - %q - %q", body.Title, body.Content, body.TopicID)
		http.Error(w, "Body Is Empty", http.StatusBadRequest)
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
			log.Println(body.TopicID)
			http.Error(w, "This ID of Topic doesn't Exist ", http.StatusBadRequest)
			return
		}
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
}
