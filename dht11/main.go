package main

import (
	"fmt"
	"github.com/d2r2/go-dht"
)

func main() {
	temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(dht.DHT11, 17, false, 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n", temperature, humidity, retried)
}
