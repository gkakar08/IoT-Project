package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)




var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}




func main() {

    // Initialize logs
    mqtt.DEBUG = log.New(os.Stdout, "\tMQTT DEBUG: ", 0)
    mqtt.ERROR = log.New(os.Stdout, "\tMQTT ERROR: ", 0)

    // Setup client options
    opts := mqtt.NewClientOptions()

    opts.AddBroker(fmt.Sprintf("%s://%s:%d", BROKER_SCHEME, BROKER_HOST, BROKER_PORT))

    opts.SetClientID(CLIENT_ID)
    opts.SetUsername(CLIENT_USER)
    opts.SetPassword(CLIENT_PSWD)

    opts.SetKeepAlive(60 * time.Second)

    // Set the message callback handler
    opts.SetDefaultPublishHandler(messageHandler)
    // opts.SetPingTimeout(30 * time.Second)

    // Initialize client
    c := mqtt.NewClient(opts)
    if token := c.Connect(); token.Wait() && token.Error() != nil {
	panic(token.Error())
    }

    // Subscribe to a topic
    if token := c.Subscribe(TOPIC, 0, nil); token.Wait() && token.Error() != nil {
	fmt.Println(token.Error())
	os.Exit(1)
    }

    time.Sleep(6 * time.Second)

    // Publish a message
    token := c.Publish(TOPIC, 0, false, "Hello from Golang!")
    token.Wait()

    time.Sleep(6 * time.Second)

    /*

    // Unscribe
    if token := c.Unsubscribe(TOPIC); token.Wait() && token.Error() != nil {
	fmt.Println(token.Error())
	os.Exit(1)
    }

    // Disconnect
    c.Disconnect(250)
    time.Sleep(1 * time.Second)
    */
    time.Sleep(60*time.Second)
}
