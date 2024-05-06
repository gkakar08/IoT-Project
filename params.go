/*
APP_ID="a953fdf7"
APP_SECRET="q7be7eda1d1d9d10"

API_ENDPOINT="https://y4d72f30.ala.us-east-1.emqxsl.com:8443/api/v5"

BROKER_ADDR="y4d72f30.ala.us-east-1.emqxsl.com"
BROKER_TLS_PORT=8883
*/

package main

const (
    BROKER_ADDR = "broker.emqx.io"
    BROKER_TCP  = 1883
    BROKER_WS   = 8083
    BROKER_TLS  = 8883
    BROKER_WSS  = 8084
    BROKER_QUIC = 14567

    CLIENT_ID = "go_mqtt_client"
    CLIENT_USER = "emqx"
    CLIENT_PSWD = "public"
)
