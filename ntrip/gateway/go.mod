// Implementation of an NTRIP Gateway for MQTT
module ntripgateway

go 1.12

require (
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/geoscienceaustralia/go-rtcm v0.0.0-20190613105737-2dc151e09ca6
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.2
	github.com/sirupsen/logrus v1.4.2
)