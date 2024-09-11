package natsutil

import (
	"github.com/nats-io/nats.go"
	"log"
)

func ConnectNATS() *nats.Conn {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	return nc
}

func PublishMessage(nc *nats.Conn, subject, message string) {
	err := nc.Publish(subject, []byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func SubscribeToSubject(nc *nats.Conn, subject string, handler nats.MsgHandler) {
	_, err := nc.Subscribe(subject, handler)
	if err != nil {
		log.Fatal(err)
	}
}
