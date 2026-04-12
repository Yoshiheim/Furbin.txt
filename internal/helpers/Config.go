package helpers

import (
	"hoxt/data"
	"net/http"
)

func GetConfig(w http.ResponseWriter, path string) {
	data.LoadConfig("./")
}
