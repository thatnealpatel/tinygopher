package main

import (
	"time"

	"github.com/thatnealpatel/tinygopher/internal/oled"
)

const sleepIvl = 3333 * time.Millisecond

func main() {
	println(oled.SetupDisplay())
	time.Sleep(50 * time.Millisecond) // todo: remove if stable

	println("text", oled.Text(11, 41, "go rsc()", 0x12B))
	time.Sleep(3333 * time.Millisecond)

	// todo: consider "Save Memory" under https://tinygo.org/docs/guides/tips-n-tricks/
	println("entering frame loop")
	for i := 0; ever(); i++ {
		x := 128
		if i > 1 {
			x = 64
		}
		println("draw[", i, "]: ", oled.Draw(x, 64, oled.AllGophers[i])) // should not copy
		time.Sleep(sleepIvl)
		if i == len(oled.AllGophers)-1 {
			i = 0
		}
	}
}

// ever is a small QoL function to ensure minimum sleep inside tight loop.
func ever() bool {
	time.Sleep(83 * time.Millisecond)
	return true
}
