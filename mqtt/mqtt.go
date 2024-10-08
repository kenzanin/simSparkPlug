package mqtt

import (
	"config"
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Connect lost: %v", err)
}

func Init(c *config.Config) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(c.Broker)
	opts.SetClientID(c.ClientId)
	opts.SetUsername(c.UserName)
	opts.SetPassword(c.Password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic(token.Error())
	}
	return client, nil
}

func Publish(client mqtt.Client, data *[]config.SparkPlugB) {
	for i := 0; i < len(*data); i++ {
		topic := (*data)[i].Topic
		payload, err := json.Marshal((*data)[i])
		if err != nil {
			log.Fatalf("error convert data to json: %s", err)
			return
		}
		token := client.Publish(topic, 0, false, payload)
		token.Wait()
	}
}
