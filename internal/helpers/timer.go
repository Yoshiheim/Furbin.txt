package helpers

import (
	"hoxt/data"
	"hoxt/internal/db"
	"hoxt/internal/modules"
	"log"
	"time"
)

var Dest time.Duration

func Timer() {
	Dest, err := ParseCustomDuration(data.Configs.ClearTimer.Temp)
	if err != nil {
		log.Println("CANT PARSE CONFIG")
		return
	}
	go func() {
		tick := time.NewTicker(Dest)
		for range tick.C {
			db.DB.Where("is_titled = ?", data.Configs.ClearTimer.ClearPinned).Delete(&modules.Paste{})
		}
	}()
}
