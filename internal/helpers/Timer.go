package helpers

import (
	"hoxt/data"
	"hoxt/internal/db"
	"hoxt/internal/modules"
	"log"
	"time"
)

var Dest time.Duration

// Delete All Pastes, even pinned, if in "/HOXT/data/config.json" change "destroy_pinned" into "true"(by default its "false") JSON config in "clear_timer".
func Timer() {
	Dest, err := ParseCustomDuration(data.Configs.ClearTimer.Temp)
	if err != nil {
		log.Fatalln("CANT PARSE CONFIG")
		return // log.Fatalln will project with code '1', so why return here lol.
	}
	go func() {
		tick := time.NewTicker(Dest)
		for range tick.C {
			db.DB.Where("is_titled = ?", data.Configs.ClearTimer.ClearPinned).Delete(&modules.Paste{})
		}
	}()
}
