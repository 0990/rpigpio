package main

import (
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
	"fmt"
)

const (
	RPin = 13
	GPin = 19
)

func main() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	pin12:=rpio.Pin(12)
	pin12.Mode(rpio.Pwm)
	pin12.Freq(64000)
	pin12.DutyCycle(255,255)

	pin := rpio.Pin(19)

	pin.Mode(rpio.Pwm)
	pin.Freq(64000)
	pin.DutyCycle(255, 255)
	time.Sleep(time.Second)
	pin.Freq(32000)
	time.Sleep(time.Second)
	fmt.Println("125")
	pin.DutyCycle(125,255)
	time.Sleep(2*time.Second)
	for i := 0; i < 5; i++ {
		for i := uint32(0); i < 255; i++ {
			pin.DutyCycle(i, 255)
			time.Sleep(time.Second / 255)
		}
		for i := uint32(255); i > 0; i-- {
			pin.DutyCycle(i, 255)
		}
	}
}

func rgb(r, g, b uint32) {

}
