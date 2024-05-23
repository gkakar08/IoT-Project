package sensor

import "time"

/** SENSOR INTERFACE (abstract class) **/

// Apart from functionality, require that each sensor has a dedicated constructor
// returning a Sensor object (or pointer), AND, a configuration table (when implemented)
type Sensor interface {

	// Name to use when identifying this **type** of sensor for the current device
	Name() string

	// The name of the units used by the sensor
	Unit() string

	// The topic name to use for the broker's identification
	Topic() string

	// Currently set measurement intervals
	// Must span at least the minimum allowed interval for the sensor
	Interval() time.Duration

	// Miniumum allowed measurement intervals
	MinInterval() time.Duration

	// Run a test measurement - return any error status, if any; nil otherwise
	MeasureStat() error

	// Measure() performs the measurement and retrieves the data value
	Measure() (float64, error)

	// Config() parses given request and integrates the configuration
	DispatchConfig(c string) error

	// ConfigHelp() returns the configuration command settings as a displayable string
	// ConfigHelp() string
}




/** CONFIGURATION MANAGEMENT (TBI) **/