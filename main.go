package main

import (
	"time"

	"github.com/thatnealpatel/tinygopher/internal/oled"
)

func main() {
	println(oled.SetupDisplay())
	time.Sleep(time.Second) // todo: remove

	frames := [][512]byte{oled.GopherBelly, oled.GopherDrink, oled.GopherGraduate}
	for {
		for i := range frames {
			oled.Draw(frames[i][:])
			time.Sleep(time.Second)
		}

		time.Sleep(83 * time.Millisecond) // always
	}
}
