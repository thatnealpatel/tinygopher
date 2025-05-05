package main

import (
	"time"

	"github.com/thatnealpatel/tinygopher/internal/oled"
)

func main() {
	println(oled.SetupDisplay())
	time.Sleep(time.Second) // todo: remove if stable

	frames := [][512]byte{oled.GopherBelly, oled.GopherDrink, oled.GopherGraduate}
	for ever() {
		for i := range frames {
			oled.Draw(frames[i][:])
			time.Sleep(time.Second)
		}
	}
}

// ever is a small QoL function to ensure minimum polling rate.
func ever() bool {
	time.Sleep(83 * time.Millisecond)
	return true
}
