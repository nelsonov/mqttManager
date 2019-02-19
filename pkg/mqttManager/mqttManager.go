package mqttManager // import "github.com/nelsonov/mqttManager"

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var broker string
var topic string
var payload []byte

func onConnect(c MQTT.Client) {
	if token := c.Subscribe(
		topic, 0, f); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
}

func messageHandler(client MQTT.Client, msg MQTT.Message) {
	payload = msg.Payload()
}
var f MQTT.MessageHandler = messageHandler

func initBroker() *MQTT.ClientOptions {
	opts := MQTT.NewClientOptions().AddBroker(broker)
	var oc MQTT.OnConnectHandler = onConnect
	opts.OnConnect = oc
	return opts
}

func connectBroker(opts *MQTT.ClientOptions) (bool, error) {
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return false, token.Error()
	} else {
		return true, nil
	}
}
