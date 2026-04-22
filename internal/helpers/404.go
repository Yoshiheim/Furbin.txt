package helpers

import (
	"html/template"
	"log"
	"net/http"
)

func Render404(w http.ResponseWriter) {
	tpl, err := template.New("404.html").Funcs(FuncMap).ParseFiles("./templates/404.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, nil); err != nil {
		log.Println(err.Error())
		http.Error(w, "Cant Parse File", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
