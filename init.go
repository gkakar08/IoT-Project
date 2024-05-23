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
	sensor := device.InitFauxSensor()
	err := pub.AddSensor(sensor)
	if err != nil {
		return err
	}

	return nil
}