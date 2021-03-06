// NTRIP Relay Adapter for MQTT
package main

import (
	"flag"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-gnss/rtcm/rtcm3"
	"github.com/go-gnss/ntrip"
	"time"
)

func main() {
	broker := flag.String("broker", "tcp://localhost:1883", "MQTT broker")
	topic := flag.String("topic", "ALIC00AUS", "MQTT topic prefix")
	source := flag.String("caster", "http://one.auscors.ga.gov.au:2101/ALIC7", "NTRIP caster mountpoint to stream from")
	username := flag.String("username", "", "NTRIP username")
	password := flag.String("password", "", "NTRIP password")
	timeout := flag.Duration("timeout", 2*time.Second, "NTRIP reconnect timeout")
	flag.Parse()

	opts := mqtt.NewClientOptions()
	opts.AddBroker(*broker)
	mqttClient := mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	ntripClient, _ := ntrip.NewClient(*source)
	ntripClient.SetBasicAuth(*username, *password)

	for ; ; time.Sleep(*timeout) {
		resp, err := ntripClient.Connect()
		if err != nil || resp.StatusCode != 200 {
			fmt.Println("NTRIP client failed to connect -", resp, err)
			continue
		}
		fmt.Println("NTRIP client connected")

		scanner := rtcm3.NewScanner(resp.Body)
		for msg, err := scanner.Next(); err == nil; msg, err = scanner.Next() {
			mqttClient.Publish(fmt.Sprintf("%s/%d", *topic, msg.Number()), 1, false, msg.Serialize())
		}

		fmt.Println("NTRIP client disconnected -", err)
	}
}
