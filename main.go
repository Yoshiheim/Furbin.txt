package main

import (
	"embed"
	"flag"
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

	hostflag := flag.String("host", data.Configs.Host, "Host Of Website")
	portflag := flag.String("port", data.Configs.Port, "Port Of Website")

	flag.Parse()

	if *hostflag != data.Configs.Host && *portflag == data.Configs.Port {

		log.Printf("Server ran on http://%s:%s\n", data.Configs.Host, data.Configs.Port)

		http.ListenAndServe(data.Configs.Host+":"+data.Configs.Port, nil)
	} else {
		log.Printf("Server ran on http://%s:%s\n", *hostflag, *portflag)

		http.ListenAndServe(*hostflag+":"+*portflag, nil)

	}
}
