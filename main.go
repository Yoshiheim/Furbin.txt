package main

import (
	"embed"
	"hoxt/data"
	"hoxt/internal/db"
	"hoxt/internal/helpers"
	"hoxt/internal/router"
	"log"
	"net/http"
)

var staticFiles embed.FS

func main() {
	db.InitDataBase()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	data.InitConfig("./data/config.json")

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, data.Configs.FaviconPath)
	})

	router.InitRoute()

	helpers.Timer()

	log.Printf("Server Runned in http://%s:%s\n", data.Configs.Host, data.Configs.Port)

	http.ListenAndServe(data.Configs.Host+":"+data.Configs.Port, nil)
}
