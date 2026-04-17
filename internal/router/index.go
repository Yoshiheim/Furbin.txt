package router

import (
	"hoxt/internal/handlers"
	"hoxt/internal/helpers"
	"net/http"
)

func InitRoute() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/topic/", handlers.FindByTopic)
	http.HandleFunc("/paste/", handlers.Local)
	http.Handle("/create", helpers.LimitMiddleware(http.HandlerFunc(handlers.CreatePaste)))
	//http.HandleFunc("/create-paste", handlers.CreatePasteHtmlTemplate)
}
