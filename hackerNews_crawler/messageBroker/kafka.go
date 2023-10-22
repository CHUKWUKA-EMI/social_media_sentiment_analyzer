package messagebroker

import (
	"context"
	"crypto/tls"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func writeMessage(msg string){
   mechanism, err := scram.Mechanism(scram.SHA512,
	os.Getenv("KAFKA_USERNAME"),
	os.Getenv("KAFKA_PASSWORD"))

   if err !=nil{
	log.Fatalln(err)
   }

	  sharedTransport := &kafka.Transport{
				SASL: mechanism,
				TLS: &tls.Config{},
			}
	writer := &kafka.Writer{
      Addr: kafka.TCP(os.Getenv("KAFKA_ENDPOINT")),
	  Transport: sharedTransport,
      Balancer: &kafka.Hash{},
	  Topic: os.Getenv("KAFKA_TOPIC"),
	}
    defer writer.Close()

	error := writer.WriteMessages(context.Background(),
	 kafka.Message{
		Value: []byte(msg),
	 },
               )
    if error !=nil{
		log.Fatalln("failed to write message:", err)
	}
}