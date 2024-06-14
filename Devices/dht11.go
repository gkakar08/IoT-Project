package device

import (
	"time"

	dht "github.com/d2r2/go-dht"
)

/** DEFINE SENSOR DATA STRUCTURE **/

// Sensor: FauxSensor
type DHT11Temp struct {
	Pin int
	Retries int
	Delay time.Duration
	MinDelay time.Duration
}

// Constructor
func InitDHT11(pin, retries int) *DHT11Temp {
	t := 1*time.Second
	return &DHT11Temp{ pin, retries, t, t }
}




/** IMPLEMENT INTERFACE METHODS **/

func (sensor *DHT11Temp) Name() string {
	return "dht11"
}

func (sensor *DHT11Temp) Unit() string {
	return "*C"
}

func (sensor *DHT11Temp) Topic() string {
	return sensor.Name()
}

func (sensor *DHT11Temp) Interval() time.Duration {
	return sensor.Delay
}

func (sensor *DHT11Temp) MinInterval() time.Duration {
	return sensor.MinDelay
}

func (sensor *DHT11Temp) MeasureStat() error {
	_, err := sensor.Measure()
	return err
}

func (sensor *DHT11Temp) Measure() (float64, error) {
	// Ignore the humidity measurement for now...
	// Need to implement dynamic measurements
	//  => standardize JSON or some other format throughout all communication (TBD)
	temperature, _, _, err := dht.ReadDHTxxWithRetry(dht.DHT11, sensor.Pin, false, sensor.Retries)
	return float64(temperature), err
}

func (sensor *DHT11Temp) DispatchConfig(c string) error {
	// No config settings yet... (TBD)
	return nil
}
