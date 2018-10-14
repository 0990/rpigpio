package main

import (
	"fmt"
	"github.com/MichaelS11/go-dht"
)

func main() {
	err := dht.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return
	}

	dht, err := dht.NewDHT("GPIO17", dht.Fahrenheit, "dht11")
	if err != nil {
		fmt.Println("NewDHT error:", err)
		return
	}

	humidity, temperature, err := dht.ReadRetry(11)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Printf("humidity:%v\n", humidity)
	fmt.Printf("temperature:%v\n", temperature)
}
