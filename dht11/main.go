package main

import (
	"fmt"
	"github.com/d2r2/go-dht"
	"log"
)

func main() {
	//err := dht.HostInit()
	//if err != nil {
	//	fmt.Println("HostInit error:", err)
	//	return
	//}
	//
	//dht, err := dht.NewDHT("GPIO17", dht.Fahrenheit, "dht11")
	//if err != nil {
	//	fmt.Println("NewDHT error:", err)
	//	return
	//}
	//
	//humidity, temperature, err := dht.ReadRetry(11)
	//if err != nil {
	//	fmt.Println("Read error:", err)
	//	return
	//}
	//
	//fmt.Printf("humidity:%v\n", humidity)
	//fmt.Printf("temperature:%v\n", temperature)
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 17, true, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)
}
