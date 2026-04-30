package helpers

import (
	"fmt"
	"hoxt/data"
	"runtime"
	"time"
)

func MemoryUsageTick() {
	Dest, err := ParseCustomDuration(data.Configs.CheckMemoryUsageTick)
	if err != nil {
		return
	}
	go func() {
		tick := time.NewTicker(Dest)
		for range tick.C {
			MemoryUsage()
		}
	}()
}

func MemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v]\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
