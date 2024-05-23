package main

/** DEVICE CONFIGURATION **/

// Hard-coded device configurations
// => ONLY change this part as the device or broker instance change
const (
	BROKER_HOST = "y4d72f30.ala.us-east-1.emqxsl.com"
	CLIENT_ID   = "Shah-Laptop"
	CLIENT_USER = "UoGuelph-IoT"
	CLIENT_PSWD = "UoGuelph-IoT-Sensors"
)




/** CONSTANTS **/

// Logging flags
const (
	DEBUG_ENABLED = true
	ERROR_ENABLED = true
)

// EMQX Broker Supported Transport Protocols
// Note: Do not use Web Sockets (wss) in Golang; incompatible. Works fine in, e.g., Python
const (
    SSL_SCHEME = "ssl"
	SSL_PORT   = 8883

	WSS_SCHEME = "wss"
	WSS_PORT   = 8084
)

// Topic sub-directory names
const (
	TOPIC_ROOT   = "root"
	TOPIC_CONFIG = "config"
	TOPIC_DATA   = "data"
)
