package main

import (
	"fmt"
	"time"

	"github.com/The-Ahmed-Shahriar/IoT-Project/sensor"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

/** PUBLISHER STRUCT (mediator pattern) **/

// Definition
type Publisher struct {
	client mqtt.Client
	sensors []sensor.Sensor
}




/** INITIALIZATION FUNCTIONS **/

// Constructor
func InitPublisher() (*Publisher, error) {

	// Initialize and connect MQTT client
	client := mqtt.NewClient(getClientOptions())
	token := client.Connect()
	if isBadToken(token, 10*time.Second) {
		return nil,token.Error()
    }

	// Initialize sensors
	sensors := make([]Sensor, 0)

	return &Publisher{ client, sensors },nil
}


// Configure MQTT client options
func getClientOptions() *mqtt.ClientOptions {

	opts := mqtt.NewClientOptions()

	// Set broker properties
    opts.AddBroker(fmt.Sprintf("%s://%s:%d", SSL_SCHEME, BROKER_HOST, SSL_PORT))

    opts.SetClientID(CLIENT_ID)
    opts.SetUsername(CLIENT_USER)
    opts.SetPassword(CLIENT_PSWD)

    opts.SetKeepAlive(60 * time.Second)
	// opts.SetPingTimeout(5 * time.Second)

    // Set default event handler(s)
    opts.SetDefaultPublishHandler(DEFAULT_MSG_HANDLER)

	return opts
}


// Add a sensor and returns any functionality errors from running the sensor
// => nil is returned if and only if sensor functionl and added to tracking list
func (pub *Publisher) AddSensor(sensor Sensor) error {

	// Issue registry if sensor is functional
	stat := sensor.MeasureStat()
	if stat != nil {
		pub.sensors = append(pub.sensors, sensor)

		// Add configuration handler for sensor through the broker's topic channel
		topic := fmt.Sprintf("%s/%s/%s", TOPIC_ROOT, sensor.Topic(), TOPIC_CONFIG)		// configuration topic
		pub.client.AddRoute(topic, makeConfigHandler(sensor))
	}
	return stat
}




/** MAIN FUNCTIONALITY **/

// Start all processes up, indefinitely
// Note: this is a BLOCKING CALL
func (pub *Publisher) Run() error {

	// Distribute publishing duties per goroutine
	for _,sensor := range pub.sensors {
		go pub.handleSensor(sensor)
	}

	// Need to implement halting/signaling channels for more complex device level control (TBI)
	for {
		// ...
	}
}




/** EVENT HANDLERS / EVENT HANDLER GENERATORS **/

// Handle individual sensors (intended to run on its own goroutine)
// Note: Implemented as a method to emphasize the need for concurrency control on accessing the client instance
func (pub *Publisher) handleSensor(sensor Sensor) {

	for {
		// Delay by set timing
		time.Sleep(sensor.Interval())

		// Fetch the measurement from sensor
		val, err := sensor.Measure()
		if err != nil {
			fmt.Printf("ERROR: Could not measure using ", sensor.Name(), " sensor")
			continue
		}

		topic := fmt.Sprintf("%s/%s/%s/%s", TOPIC_ROOT, sensor.Topic(), TOPIC_DATA, CLIENT_ID)
		token := pub.client.Publish(topic, 0, true, val)
		if isBadToken(token, 1*time.Second) {
			fmt.Printf("ERROR: Message failed to acknowledge for ", sensor.Name(), "'s measurement")
		}
	}
}


// 
func makeConfigHandler(sensor Sensor) mqtt.MessageHandler {

	handler := func(client mqtt.Client, msg mqtt.Message) {

			err := sensor.DispatchConfig(string(msg.Payload()))
			if err != nil {
				fmt.Println("CONFIG ERROR: ", err)
			}
			
			// if error, send back a status + config help message (TBI)
			/*
			if err != nil {
				client.Publish(fmt.Sprintf("root/%s/configRes/"))
			}
			*/
		}
	
	return handler
}


// Basic handler for general messages
var DEFAULT_MSG_HANDLER mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}


// Determine a bad request given the request token
func isBadToken(token mqtt.Token, timeout time.Duration) bool {
	return token.WaitTimeout(timeout) && token.Error() != nil
}
