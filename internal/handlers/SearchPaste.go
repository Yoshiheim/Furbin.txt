package handlers

import (
	"fmt"
	"hoxt/internal/db"
	"hoxt/internal/helpers"
	"hoxt/internal/modules"
	"html/template"
	"log"
	"net/http"
)

type NewPaste struct {
	ID    uint
	Title string
}

func SearchPaste(w http.ResponseWriter, r *http.Request) {

	keyword := r.FormValue("keyword")

	keyword = helpers.DestroySpaces(keyword)

	keyword = helpers.OnlyASCII(keyword)

	keyword = helpers.TrimLeft(keyword)

	if helpers.CheckSizeString(keyword, 35) {
		fmt.Fprintf(w, "Bro this keyword is sooo big(35 symbols limit)")
		return
	}

	id, preid, nextid, err := helpers.SafeParsePage(r)
	if err != nil {
		http.Error(w, "error wit args", http.StatusBadRequest)
		return
	}

	var pastes []modules.Paste

	page := int(id)
	limit := 10
	offset := (page - 1) * limit

	if offset < 0 {
		offset = 0
	}

	if page < 1 {
		page = 1
	}

	if act := db.DB.
		Order("id ASC").
		Offset(offset).
		Where("title LIKE ?", "%"+keyword+"%").
		Limit(limit).
		Find(&pastes); act.Error != nil {
		fmt.Println(act.Error.Error())
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}

	tpl, err := template.New("SearchPaste.html").Funcs(helpers.FuncMap).ParseFiles("./templates/SearchPaste.html", "./templates/search.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}

	//and render for client
	tpl.Execute(w, map[string]any{
		"pastes":  pastes,
		"keyword": keyword,
		"preid":   preid,
		"nextid":  nextid,
	})
}
