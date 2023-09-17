package main

import (
	"log"
	"github.com/nats-io/nats.go"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the NATS server, use docker network DNS name
	natsURL := "nats://nats:4222"
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	//SERVER MODE -> produce message when you contact the root
	
	r := gin.Default()

	// http://localhost:8080/publish?msg=????
	r.GET("/publish", func(c *gin.Context) {				
		subject := "NATS_sub"
		message := c.Query("msg")

		err = nc.Publish(subject, []byte(message))
		if err != nil {
			log.Fatal(err)
			c.String(500, "NATS not reachable")
		}else {
			c.String(200, message)
		}

	})

	// Start the server on port 8080
	r.Run(":8080")



	//AUTOMATIC MODE -> automatically produce messages
	// Subject to which we'll publish the message
	
	/*
	subject := "NATS_sub"

	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Second)

		// Message to send
		message := fmt.Sprintf("NATS message :%d",i)

		// Publish the message
		err = nc.Publish(subject, []byte(message))
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Published: %s", message)
	}
	*/
	
}

