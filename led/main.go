package main

import (
	"fmt"
	"github.com/nathan-osman/go-rpigpio"
	"os"
	"os/signal"
	"time"
)

func main() {
	p, err := rpi.OpenPin(17, rpi.OUT)
	if err != nil {
		panic(err)
	}
	defer func() {
		p.Close()
	}()

	go func() {
		for {
			fmt.Println("output high")
			p.Write(rpi.HIGH)

			time.Sleep(time.Second * 1)
			fmt.Println("output low")
			p.Write(rpi.LOW)
			time.Sleep(time.Second * 1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c)
	s := <-c
	fmt.Println("Got signal:", s)
}
