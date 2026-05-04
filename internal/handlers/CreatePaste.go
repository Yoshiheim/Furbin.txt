package handlers

import (
	"encoding/json"
	"fmt"
	"hoxt/data"
	"hoxt/internal/db"
	"hoxt/internal/helpers"
	"hoxt/internal/modules"
	"html"
	"net/http"
	"unicode/utf8"
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
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⣇⠀⠀⠀⠀⠀⢸⡇⠀⠀⠀⠀⠀⠀⣿⣾⡽⠫⣱⣷⠝⠘⣧⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠀⠀⠀⠀⠀⠀⣇⠀⠀⠀⠀⠀⠀⢹⠀⠀⠀⠀⠀⠀⠀⢸⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⠀⠀⠀⠀⠈⢷⠀⠀⠀⠀⠀⠀⢹⣧⠑⢧⣻⢄⠣⢿⣧⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢼⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⠀⠀⠀⠀⠀⠀⢻⠂⠀⠀⠀⠀⠰⡜⠀⠀⠀⠀⠀⠀⠀⢺⡂
- YOU CAN POST PASTES WITH ASCII ART LIKE THIS BOYKISSER.
*/

// Create Paste in Topic as JSON Post Request.
// path: 'http://<HOST>:<PORT>/create'
func CreatePaste(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 280*1024)

	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Author  string `json:"author"`
	}

	// decode user's request
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Cannot parse JSON body", http.StatusBadRequest)
		return
	}

	// data.Configs.PasteLens.AuthorLen its from "/HOXT/data/config.json", in the "paste_lens" {"title_len"}
	if utf8.RuneCountInString(body.Title) > data.Configs.PasteLens.TitleLen {
		http.Error(w, "Title text-field exceeds character limit of 128.", http.StatusBadRequest)
		return
	}

	// data.Configs.PasteLens.ContentLen its from "/HOXT/data/config.json", in the "paste_lens" {"content_len"}
	// (65535 = 64kb) btw
	if utf8.RuneCountInString(body.Content) > data.Configs.PasteLens.ContentLen {
		http.Error(w, "Content text-field exceeds character limit of 65536.", http.StatusBadRequest)
		return
	}

	// data.Configs.PasteLens.ContentLen its from "/HOXT/data/config.json" to json: "paste_lens" {"author_len"}
	if utf8.RuneCountInString(body.Author) > data.Configs.PasteLens.AuthorLen {
		http.Error(w, "Author text-field exceeds character limit of 128.", http.StatusBadRequest)
		return
	}

	// escape all content
	body.Title = html.EscapeString(helpers.OnlyASCII(helpers.TruncateByte(helpers.DestroySpaces(body.Title), data.Configs.PasteLens.TitleLen)))
	body.Content = html.EscapeString(helpers.OnlyASCII(helpers.TruncateByte(body.Content, data.Configs.PasteLens.ContentLen)))
	body.Author = html.EscapeString(helpers.OnlyASCII(helpers.TruncateByte(helpers.DestroySpaces(body.Author), data.Configs.PasteLens.AuthorLen)))

	//Check is 'title' in JSON requet is empty.
	if helpers.DestroySpaces(body.Title) == "" {
		http.Error(w, "Title Is empty", http.StatusBadRequest)
		return
	}

	//same but with 'content'.
	if helpers.DestroySpaces(body.Content) == "" {
		http.Error(w, "Content Is empty", http.StatusBadRequest)
		return
	}

	fmt.Printf("[%s], [%s]\n", body.Title, body.Content)

	// Create Paste On DB.
	// 'author' in JSON request is optional btw.
	paste, err := helpers.CreatePasteIfTopicExists(db.DB, modules.Paste{
		Title:   body.Title,
		Content: body.Content,
		Author:  body.Author,
	})
	if err != nil {
		http.Error(w, "we have some problem with database.", http.StatusInternalServerError)
		return
	}

	// Create Paste On DB.
	// We dont need this code.
	/*	act := db.DB.Create(&modules.Paste{
				Title:   body.Title,
				Content: body.Content,
				Author:  body.Author,
				TopicID: body.TopicID,
			})

		// If DB Query have Error, Check kind of Error, otherwise http.StatusInternalServerError Idk Why.
		if act.Error() != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return

		}
	*/
	// helpers.EncodeJson(w, map[string]interface{}{
	// 	"paste": paste,
	// })
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paste.ID)
}
