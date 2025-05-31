package main

import (
	"time"

	"github.com/thatnealpatel/tinygopher/internal/oled"
)

const sleepIvl = 3333 * time.Millisecond

func main() {
	println(oled.SetupDisplay())
	time.Sleep(50 * time.Millisecond)

	oled.Text(11, 41, "tinygopher ", 0x12B)
	time.Sleep(4000 * time.Millisecond)

	// todo: consider "Save Memory" under
	// https://tinygo.org/docs/guides/tips-n-tricks/
	println("entering frame loop")
	for i := 0; ever(); i++ {
		x := 128
		if i > 1 {
			x = 64
		}
		oled.Draw(x, 64, oled.AllGophers[i])
		time.Sleep(sleepIvl)
		if i == len(oled.AllGophers)-1 {
			i = -1 // elides modulo and large ints
		}
	}
}

// ever is a vanity function
// that sleeps a fixed duration.
func ever() bool {
	time.Sleep(83 * time.Millisecond)
	return true
}
