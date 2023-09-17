package main

import (
	"log"
	"time"
	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to the NATS server, use docker network DNS name
	natsURL := "nats://nats:4222" 
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subject to subscribe to
	subject := "NATS_sub"

	// Subscribe to the subject and read the messages when up
	_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
		time.Sleep(2 * time.Second)
		log.Printf("Received: %s", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Subscriber connected and listening...")

	// Keep the program running
	select {}
}
