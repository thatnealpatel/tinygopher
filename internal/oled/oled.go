package oled

import (
	"fmt"
	m "machine"
	"sync"

	"tinygo.org/x/drivers/ssd1306"
)

const (
	// (unclear) https://www.espressif.com/sites/default/files/documentation/0a-esp8266ex_datasheet_en.pdf
	// Since it looks to me that datasheet indicates MTMS + GPIO2 should be used (in the I2C section);
	// however, Pin4, Pin5 is correct: https://docs.micropython.org/en/latest/esp8266/tutorial/ssd1306.html
	sdaPin = m.GP4 // data/command
	sclPin = m.GP5 // reset
)

var (
	once    sync.Once
	display = ssd1306.Device{} // technically NOT thread-safe
)

// SetupDisplay safely initializes a global display for use.
func SetupDisplay() error {
	i2c, err := setupI2C(sdaPin, sclPin)
	if err != nil {
		return err
	}
	once.Do(func() {
		display = ssd1306.NewI2C(i2c)
		display.Configure(ssd1306.Config{
			Width:   128,
			Height:  64,
			Address: 0x3C,
		})
		display.ClearDisplay()
	})

	return nil
}

// Draw provides a semi-safe (if used as a singleton in TinyGo) drawing wrapper.
func Draw(buf []byte) (err error) {
	display.ClearDisplay()
	err = display.SetBuffer(buf)
	display.Display()
	return
}

func glide(buf []byte, frames int) (err error) {
	return
}

// setupI2C defines a software I2C implementation over the ESP8266EX designated pins.
func setupI2C(sda m.Pin, scl m.Pin) (*m.I2C, error) {
	c := m.I2CConfig{
		Frequency: m.TWI_FREQ_400KHZ, // I2C on ESP8266EX supports up to 100 kHz
		SDA:       sda,
		SCL:       scl,
	}
	if err := m.I2C0.Configure(c); err != nil {
		return m.I2C0, fmt.Errorf("cannot configure(%+v) software i2c: %v\n", c, err)
	}
	return m.I2C0, nil
}
