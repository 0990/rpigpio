package main

import (
	"fmt"
	"github.com/nathan-osman/go-rpigpio"
	"time"
)

func main() {
	p, err := rpi.OpenPin(17, rpi.OUT)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	fmt.Println("output high")
	p.Write(rpi.HIGH)

	time.Sleep(time.Second * 1)
	fmt.Println("output low")
	p.Write(rpi.LOW)

	fmt.Println("exit")
}
