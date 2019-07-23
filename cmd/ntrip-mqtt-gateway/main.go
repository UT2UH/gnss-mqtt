package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.ErrorLevel)

	// TODO: Define from config file
	gateway, err := NewGateway("2101", "tcp://vernemq:1883")
	if err != nil {
		log.Fatal(err)
	}

	gateway.Hostname = "ntrip.geops.team"
	gateway.Operator = "Geoscience Australia"

	log.Fatal(gateway.Serve())
}
