package main

import (
    "fmt"
    mqtt "github.com/eclipse/paho.mqtt.golang"
    "time"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
    fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    fmt.Printf("Connect lost: %v", err)
}

func publish(client mqtt.Client) {
    num := 10
    for i := 0; i < num; i++ {
        text := fmt.Sprintf("Message %d", i)
        token := client.Publish(TOPIC, 0, false, text)
        token.Wait()
        time.Sleep(time.Second)
    }
}

func sub(client mqtt.Client) {
    token := client.Subscribe(TOPIC, 1, nil)
    token.Wait()
    fmt.Printf("Subscribed to topic: %s", TOPIC)
}




func main() {
    opts := mqtt.NewClientOptions()

    opts.AddBroker(fmt.Sprintf("%s://%s:%d", BROKER_L4, BROKER_ADDR, BROKER_PORT))
    opts.SetClientID(CLIENT_ID)
    opts.SetUsername(CLIENT_USER)
    opts.SetPassword(CLIENT_PSWD)

    opts.SetDefaultPublishHandler(messagePubHandler)
    opts.OnConnect = connectHandler
    opts.OnConnectionLost = connectLostHandler

    fmt.Println("Client options configured")

    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
	panic(token.Error())
    }

    fmt.Println("Client initialized")

    sub(client)
    fmt.Println("Client subscribed")

    publish(client)
    fmt.Println("Client published")

    client.Disconnect(15000)
}

