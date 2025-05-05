package main

import (
	"time"

	"github.com/thatnealpatel/tinygopher/internal/oled"
)

func main() {
	println(oled.SetupDisplay())
	time.Sleep(time.Second) // todo: remove if stable

	// todo: consider "Save Memory" under https://tinygo.org/docs/guides/tips-n-tricks/
	frames := [][]byte{oled.GopherBelly, oled.GopherDrink, oled.GopherGraduate}
	for ever() {
		for i := range frames {
			oled.Draw(frames[i]) // should not copy
			time.Sleep(time.Second)
		}
	}
}

// ever is a small QoL function to ensure minimum sleep inside tight loop.
func ever() bool {
	time.Sleep(83 * time.Millisecond)
	return true
}
