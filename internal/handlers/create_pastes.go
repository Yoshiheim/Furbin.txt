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

⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣥⠙⢦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⠾⠋⠉⠁⠀⠀⠀⢠⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡇⠀⠀⠈⠛⢦⡀⠀⠀⠀⠀⠀⠀⢸⡍⠓⢦⣄⣀⠀⠀⠀⠀⠀⠀⠀⣠⠞⠁⠀⠀⠀⠀⠀⠀⠀⢸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⠁⠀⠀⠀⠀⠀⠹⢦⠀⠀⠀⠀⠀⢀⢳⣄⠀⠀⠉⠉⠓⠦⣄⠀⢠⠞⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣽⠀⠀⠀⠀⠀⠀⠀⠈⢳⣴⡖⠛⠉⠉⠉⠉⠀⠀⠀⠀⠀⠀⠙⢩⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢾⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠓⠦⣤⣀⠀⠀⠀⠀⢠⠄⠀⠀⠈⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠛⠛⠃⢀⡟⣠⡤⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠛⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣧⡀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⣠⠞⠛⠓⢦⣄⠀⠀⠀⠀⠀⠀⠀⣰⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢷⠀⠀⠀⠀⠀⠀⢠⡞⠁⠀⢀⣤⡀⠀⠀⠀⠀⠀⠈⢃⣷⣶⣆⠀⠙⣦⠀⠀⠀⢀⣠⠞⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠳⣄⢀⠀⠀⢰⠏⠀⠀⢠⣾⣿⣿⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⣇⠀⠸⡆⠀⠀⠘⣡⠴⢚⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⡄⢸⡀⠀⡟⠀⠀⠀⢸⣿⣿⣿⠀⠀⠀⠀⠀⠀⣻⣿⣿⣿⣿⠀⠀⡏⠀⠀⠀⠉⢀⡾⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⣶⡛⠋⠁⠀⢸⠀⢀⡇⠀⠀⠀⠈⢿⣿⡿⠀⠀⠀⠀⠀⠀⠈⠛⠿⠟⠁⠀⣠⣗⠀⠀⠀⣴⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⣤⡀⠀⣿⠀⠐⢳⣄⡀⡀⡀⠰⠿⠚⠛⠀⠀⠀⠀⠀⠀⠀⠠⣤⣤⡞⣱⠋⠀⠀⠀⠈⢷⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⡇⠀⢻⣄⠰⠯⢞⠙⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠑⠋⠀⠀⢀⣀⣀⣈⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⠥⢤⣀⡝⢷⣄⠈⠋⠀⠀⠀⠀⠀⠀⣴⣄⡀⠀⠀⠀⠀⠀⠀⠀⣀⣤⠶⠋⠀⠉⠈⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠓⠲⢶⣶⢤⠀⠀⠁⠀⠁⠀⠀⠀⠀⠀⠀⠀⢉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠰⢴⣞⣋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⣽⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⢃⣀⣠⡄⠀⠀⠀⠀⠀⠀⠀⢠⠀⠀⠀⠀⢹⡅⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣟⢶⣤⡀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⣸⠀⠀⠀⠀⠀⠀⠀⠀⢸⡇⠀⠀⠀⠀⢻⡀⠀⠀⠀⠀⠀⠀⣿⠵⠰⣫⠍⡌⣷⣆⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⠀⠀⠀⠀⠈⢷⠀⠀⠀⠀⠀⠀⢹⣧⠑⢧⣻⢄⠣⢿⣧⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⣇⠀⠀⠀⠀⠀⢸⡇⠀⠀⠀⠀⠀⠀⣿⣾⡽⠫⣱⣷⠝⠘⣧⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠀⠀⠀⠀⠀⠀⣇⠀⠀⠀⠀⠀⠀⢹⠀⠀⠀⠀⠀⠀⠀⢸⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢼⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⠀⠀⠀⠀⠀⠀⢻⠂⠀⠀⠀⠀⠰⡜⠀⠀⠀⠀⠀⠀⠀⢺⡂
- YOU CAN POST PASTES WITH ASCII ART LIKE THIS BOYKISSER.
*/

// Create Paste in Topic as JSON Post Request.
// path: 'http://<HOST>:<PORT>/create'
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

	//Check is 'title' in JSON requet is empty.
	if body.Title == "" {
		http.Error(w, "Title Is empty", http.StatusBadRequest)
		return
	}

	//same but with 'content'.
	if body.Content == "" {
		http.Error(w, "Content Is empty", http.StatusBadRequest)
		return
	}

	// 'author' in JSON request is optional btw.

	// Create Paste On DB.
	act := db.DB.Create(&modules.Paste{
		Title:   body.Title,
		Content: body.Content,
		Author:  body.Author,
		TopicID: body.TopicID,
	})

	// If DB Query have Error, Check kind of Error, otherwise http.StatusInternalServerError Idk Why.
	if act.Error != nil {
		if strings.Contains(act.Error.Error(), "violates foreign key constraint") {
			http.Error(w, "Topic ID does not exist ", http.StatusBadRequest)
			return
		}

		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// if all goes well, return code 200(aka http.StatusOK).
	w.WriteHeader(http.StatusOK)
}
