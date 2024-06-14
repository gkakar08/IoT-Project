package publisher

import (
	"fmt"
	"time"

	device "iot/Devices"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

/** PUBLISHER STRUCT (mediator pattern) **/

// Definition
type Publisher struct {
	client  mqtt.Client
	sensors []device.Sensor
}

/** INITIALIZATION FUNCTIONS **/

// Constructor method for a *Publisher instance
func InitPublisher() (*Publisher, error) {

	// Initialize and connect MQTT client
	client := mqtt.NewClient(getClientOptions())
	token := client.Connect()
	if token.WaitTimeout(10*time.Second) && token.Error() != nil {
		return nil, token.Error()
	}

	// Initialize sensors
	sens := []device.Sensor{}

	return &Publisher{client, sens}, nil
}

// Fetch the set MQTT client options
func getClientOptions() *mqtt.ClientOptions {

	opts := mqtt.NewClientOptions()

	// Set broker properties
	opts.AddBroker(fmt.Sprintf("%s://%s:%d", SSL_SCHEME, BROKER_HOST, SSL_PORT))
	opts.SetKeepAlive(60 * time.Second)
	// opts.SetPingTimeout(5 * time.Second)

	// Set client
	opts.SetClientID(CLIENT_ID)
	opts.SetUsername(CLIENT_USER)
	opts.SetPassword(CLIENT_PSWD)

	// Set default event handler(s)
	opts.SetDefaultPublishHandler(DEFAULT_MSG_HANDLER)

	return opts
}

// Add a sensor to the publisher's registry
// Returns nil is returned if and only if sensor functional and successfully registered
func (pub *Publisher) AddSensor(sensor device.Sensor) error {

	// Issue registry if sensor is functional
	err := sensor.MeasureStat()
	if err == nil {

		// Register the sensor
		pub.sensors = append(pub.sensors, sensor)

		// Create configuration handler
		handler := func(client mqtt.Client, msg mqtt.Message) {

			// Dispatch configuration change
			err := sensor.DispatchConfig(string(msg.Payload()))
			if err != nil {
				fmt.Println(err)
			}

			// if config request is invalid, send back a help message (TBD)
		}

		// Subscribe to the config topic with the handler
		topic := fmt.Sprintf("%s/%s/%s", TOPIC_ROOT, sensor.Topic(), TOPIC_CONFIG) // configuration topic
		token := pub.client.Subscribe(topic, 0, handler)
		err = token.Error()
	}
	return err
}

/** MAIN FUNCTIONALITY **/

// Start and maintain all processes
// Note: this is a BLOCKING CALL
func (pub *Publisher) Run() {

	done := make(chan bool)

	// Distribute publishing duties per goroutine
	for i := 0; i < len(pub.sensors); i++ {
		go pub.handleSensor(pub.sensors[i])
	}

	// Need to implement halting/signaling channels for more complex device level control (TBD)
	<-done
}

/** EVENT HANDLERS / EVENT HANDLER GENERATORS **/

// Handle individual sensors (intended to run on its own goroutine)
// Note: Implemented as a method to emphasize the need for concurrency control on accessing the client instance
func (pub *Publisher) handleSensor(sensor device.Sensor) {

	for {
		// Delay by set timing
		time.Sleep(sensor.Interval())

		// Fetch the measurement from sensor
		data, err := sensor.Measure()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Format data and determine publishing topic repository
		topic := fmt.Sprintf("%s/%s/%s/%s", TOPIC_ROOT, sensor.Topic(), TOPIC_DATA, CLIENT_ID)
		dataStr := fmt.Sprintf("%f", data)

		// Publish
		token := pub.client.Publish(topic, 0, false, dataStr)
		if token.Error() != nil {
			fmt.Println(token.Error())
		}
	}
}

// Default handler for general (unexpected) messages
var DEFAULT_MSG_HANDLER mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
