package main

import (
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
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

	pin := rpio.Pin(13)

	pin.Mode(rpio.Pwm)
	pin.Freq(64000)
	pin.DutyCycle(0, 255)
	time.Sleep(time.Second * 2)
	pin.DutyCycle(100, 255)
	time.Sleep(time.Second * 2)

	//for i := 0; i < 5; i++ {
	//	for i := uint32(0); i < 255; i++ {
	//		pin.DutyCycle(i, 255)
	//		time.Sleep(time.Second / 255)
	//	}
	//	for i := uint32(255); i > 0; i-- {
	//		pin.DutyCycle(i, 255)
	//	}
	//}
}

func rgb(r, g, b uint32) {

}
