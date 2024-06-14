package publisher

/** DEVICE CONFIGURATION **/

// Hard-coded device configurations
// => ONLY change this part as the device or broker instance change
const (
<<<<<<< HEAD
	BROKER_HOST = "g332f11e.ala.eu-central-1.emqxsl.com"
	CLIENT_ID   = "fayaaz-Latitude-5430"
	CLIENT_USER = "Fayaaz"
	CLIENT_PSWD = "d+HUi6)b!.LE_KC"
=======
	BROKER_HOST = ""
	CLIENT_ID   = ""
	CLIENT_USER = ""
	CLIENT_PSWD = ""
>>>>>>> 3b11dc6a7017b1981c1019815186a3c8c746cd48
)

/** CONSTANTS **/

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
