package oled

import (
	"fmt"
	"machine"
	"sync"

	"tinygo.org/x/drivers/ssd1306"
)

const (
	// (unclear) https://www.espressif.com/sites/default/files/documentation/0a-esp8266ex_datasheet_en.pdf
	// Since it looks to me that datasheet indicates MTMS + GPIO2 should be used (in the I2C section);
	// however, Pin4, Pin5 is correct: https://docs.micropython.org/en/latest/esp8266/tutorial/ssd1306.html
	sdaPin = machine.GP4 // data/command
	sclPin = machine.GP5 // reset
)

var (
	once    sync.Once
	display *ssd1306.Device // technically NOT thread-safe
)

// SetupDisplay safely initializes a global display for use.
func SetupDisplay() error {
	i2c, err := setupI2C(sdaPin, sclPin)
	if err != nil {
		return err
	}
	once.Do(func() {
		display = &ssd1306.NewI2C(i2c)
		display.Configure(ssd1306.Config{
			Width:   128,
			Height:  64,
			Address: 0x3C, // default
		})
		display.ClearDisplay() // in-case of weird reboot
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

// setupI2C defines a software I2C implementation over the ESP8266EX designated pins.
func setupI2C(sda, scl machine.Pin) (*machine.I2C, error) {
	c := machine.I2CConfig{
		Frequency: machine.TWI_FREQ_100KHZ, // I2C on ESP8266EX supports up to 100 kHz
		SDA:       sda,
		SCL:       scl,
	}
	if err := machine.I2C0.Configure(c); err != nil {
		return nil, fmt.Errorf("cannot configure(%+v) software i2c: %v\n", c, err)
	}
	return machine.I2C0, nil
}
