package main

import (
	"fmt"
	"log"
	"os"

	emqx "iot/EMQX"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)


func main() {

	// Setup MQTT logging
	setupLogs()

	// Initialize EMQX publisher instance
	pub, err := emqx.InitPublisher()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add all sensors **(see init.go)**
	err = setupSensors(pub)
	if err != nil {
		fmt.Println(err)
		return
	}

	pub.Run()	// This is a blocking call!
}


func setupLogs() {
    // mqtt.DEBUG = log.New(os.Stdout, "\tMQTT DEBUG: ", 0)
	mqtt.ERROR = log.New(os.Stdout, "=============\nMQTT ERROR: ", 0)
}
