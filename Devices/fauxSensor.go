package device

import (
	"fmt"
	"strings"
	"time"
)

/** DEFINE SENSOR DATA STRUCTURE **/

// Sensor: FauxSensor
type FauxSensor struct {
	Cur int						// Current call count
	Delay time.Duration
	MinDelay time.Duration
}

// Constructor
func InitFauxSensor() *FauxSensor {
	t := 1*time.Second
	return &FauxSensor{ 0, t, t }
}




/** IMPLEMENT INTERFACE METHODS **/

func (sensor *FauxSensor) Name() string {
	return "faux"
}


func (sensor *FauxSensor) Unit() string {
	return ""
}

func (sensor *FauxSensor) Topic() string {
	return sensor.Name()
}

func (sensor *FauxSensor) Interval() time.Duration {
	return sensor.Delay
}

func (sensor *FauxSensor) MinInterval() time.Duration {
	return sensor.MinDelay
}

func (sensor *FauxSensor) MeasureStat() error {
	return nil
}

func (sensor *FauxSensor) Measure() (float64, error) {
	sensor.Cur++
	return float64(sensor.Cur),nil
}

func (sensor *FauxSensor) DispatchConfig(c string) error {
	// Ideally, parse the config request as a JSON object (TBD)
	fmt.Println("Command issued: ", c)
	switch (strings.ToLower(c)) {
		case "reset":
			sensor.Cur = 0
	}
	return nil
}