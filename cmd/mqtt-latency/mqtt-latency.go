package main

import (
	"flag"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-gnss/rtcm/rtcm3"
	"time"
)

func main() {
	broker := flag.String("broker", "tcp://localhost:1883", "")
	topic := flag.String("topic", "#", "")
	flag.Parse()

	opts := mqtt.NewClientOptions().
		AddBroker(*broker).
		SetOnConnectHandler(func(client mqtt.Client) {
			if token := client.Subscribe(*topic, 1, nil); token.Wait() && token.Error() != nil {
				panic(token.Error())
			}
		})

	opts.SetDefaultPublishHandler(func(client mqtt.Client, mqttmsg mqtt.Message) {
		msg := rtcm3.DeserializeMessage(mqttmsg.Payload()) // DeserializeMessage should be able to return error
		log := fmt.Sprintf("%v %v %v %v", time.Now().Format("2006-01-02T15:04:05.999999"), mqttmsg.Topic(), msg.Number(), len(mqttmsg.Payload()))

		if obs, ok := msg.(rtcm3.Observable); ok {
			fmt.Println(log, time.Now().UTC().Sub(obs.Time()))
		} else {
			fmt.Println(log)
		}
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	select {}
}
