package oled

import (
	"errors"
	"image/color"
	"machine"

	"tinygo.org/x/drivers/i2csoft"
	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"

	"tinygo.org/x/tinyfont/freemono"
)

var (
	// sync.Once is expensive on microcontrollers.
	display ssd1306.Device // NOT thread-safe
)

// SetupDisplay initializes the global display
// device and defines the software I2C channel
// without using sychronization.
func SetupDisplay() error {
	i2c, err := setupI2C()
	if err != nil {
		return err
	}
	display = ssd1306.NewI2C(i2c)
	display.Configure(ssd1306.Config{
		Address: 0x3C, // default
		Width:   128,
		Height:  64,
	})
	display.ClearDisplay() // in-case of weird reboot
	return nil
}

// Text renders the given message at the
// provided coordinates.
func Text(x, y int16, msg string, size int) (err error) {
	// todo: support 'size' parameter
	display.ClearDisplay()
	tinyfont.WriteLine(&display, &freemono.Oblique12pt7b, x, y, msg, color.RGBA{1, 1, 1, 1})
	display.Display()
	return
}

// Draw centers the given buffer and
// renders it using no blit algorithm.
func Draw(x, y int, buf []byte) (err error) {
	// todo: generalize the centering behavior
	display.ClearDisplay()
	var ofx int16 = 0
	if x != 128 {
		ofx = 33 // "centers"
	}
	err = display.DrawBitmap(ofx, 0, pixel.NewImageFromBytes[pixel.Monochrome](x, y, buf))
	if derr := display.Display(); derr != nil {
		err = errors.New("two errors, draw(" + err.Error() + "), display(" + derr.Error() + ")")
	}
	return
}

// setupI2C defines a software I2C channel
// over the ESP8266EX designated pins.
func setupI2C() (*i2csoft.I2C, error) {
	i2c := i2csoft.New(machine.SCL_PIN, machine.SDA_PIN) // scl, sda machine.Pin
	c := i2csoft.I2CConfig{Frequency: 100_000, SCL: machine.SCL_PIN, SDA: machine.SDA_PIN}
	if err := i2c.Configure(c); err != nil {
		return nil, errors.New("cannot software i2c: " + err.Error())
	}
	return i2c, nil
}
