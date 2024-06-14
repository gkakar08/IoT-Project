package main

import (
	device "iot/Devices"
	emqx "iot/EMQX"
)

// Initialize your sensors under this function
// Returns any error status if any (nil otherwise)
//
// The following outlines a general structure for adding each of your sensors:
//
// 	- <call your sensor's constructor>
// 	- <check for error from constructor, if necessary>
// 	- err := pub.AddSensor(<your sensor variable>`)
// 	- if err != nil { return err }
//
func setupSensors(pub *emqx.Publisher) error {

	// (1) Add test/faux sensor
	sensorFaux := device.InitFauxSensor()
	err := pub.AddSensor(sensorFaux)
	if err != nil {
		return err
	}

	// (2) DHT11 sensor (temperature only)
	/*
	sensorDHT11 := device.InitDHT11(4, 10)
	err = pub.AddSensor(sensorDHT11)
	if err != nil {
		return err
	}
	*/

	return nil
}
