package helpers

import (
	"fmt"
	"hoxt/data"
	"hoxt/internal/db"
	"hoxt/internal/modules"
	"time"
)

var Dest time.Duration

// Delete All Pastes, even pinned, if in "/HOXT/data/config.json" change "destroy_pinned" into "true"(by default its "false") JSON config in "clear_timer".
func Timer() {
	Dest, err := ParseCustomDuration(data.Configs.ClearTimer.Temp)
	if err != nil {
		return
	}
	fmt.Println("[INIT TIME CLEARER]")
	go func() {
		tick := time.NewTicker(Dest)
		for range tick.C {
			db.DB.Where("is_titled = ?", data.Configs.ClearTimer.ClearPinned).Delete(&modules.Paste{})
			db.DB.Raw(`DELETE FROM sqlite_sequence WHERE name = "pastes"`)
		}
	}()
}
