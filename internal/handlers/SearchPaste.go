package handlers

import (
	"fmt"
	"hoxt/internal/db"
	"hoxt/internal/helpers"
	"hoxt/internal/modules"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type NewPaste struct {
	ID    uint
	Title string
}

func SearchPaste(w http.ResponseWriter, r *http.Request) {

	keyword := r.FormValue("keyword")

	keyword = helpers.DestroySpaces(keyword)

	keyword = helpers.OnlyASCII(keyword)

	if helpers.CheckSizeString(keyword, 35) {
		fmt.Fprintf(w, "Bro this keyword is sooo big(35 symbols limit)")
		return
	}

	var err error

	var id uint32

	var preid uint32

	var nextid uint32

	id, preid, nextid, err = helpers.SafeParsePage(r)
	if err != nil {
		http.Error(w, "error wit args", http.StatusBadRequest)
		return
	}

	var pastes []modules.Paste
	var finded_pastes []NewPaste

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
		Order("id DESC").
		Offset(offset).
		Limit(limit).
		Where("title LIKE ?", "%"+keyword+"%").
		Find(&pastes); act.Error != nil {
		fmt.Println(act.Error.Error())
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}

	if keyword != "" {
		for _, v := range pastes {
			if strings.Contains(v.Title, keyword) {
				newtitle := strings.ReplaceAll(v.Title, keyword, strings.ToUpper(keyword))
				finded_pastes = append(finded_pastes, NewPaste{
					Title: newtitle,
					ID:    v.ID,
				})
			}
			fmt.Printf("[%s]\n", v.Title)
		}
	}

	tpl, err := template.New("SearchPaste.html").Funcs(helpers.FuncMap).ParseFiles("./templates/SearchPaste.html", "./templates/search.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}

	//and render for client
	tpl.Execute(w, map[string]any{
		"pastes":  finded_pastes,
		"keyword": keyword,
		"preid":   preid,
		"nextid":  nextid,
	})
}
