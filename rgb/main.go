package main

import (
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
)

func main() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	pin := rpio.Pin(13)

	pin.Mode(rpio.Pwm)
	pin.Freq(64000)
	pin.DutyCycle(0, 32)

	for i := 0; i < 5; i++ {
		for i := uint32(0); i < 32; i++ {
			pin.DutyCycle(i, 32)
			time.Sleep(time.Second / 32)
		}
		for i := uint32(32); i > 0; i-- {
			pin.DutyCycle(i, 32)
		}
	}
}
