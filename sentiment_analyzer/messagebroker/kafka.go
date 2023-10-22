package messagebroker

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/chukwuka-emi/hackerNews_comments_sentiments/sentiment_analyzer/sentiments"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type comment struct {
	PostURL string
	Author  string `selector:"a.hnuser"`
	URL     string `selector:".age a[href]" attr:"href"`
	Comment string `selector:".comment"`
}

func readMessage(){
   mechanism, err := scram.Mechanism(scram.SHA512,
	os.Getenv("KAFKA_USERNAME"),
	os.Getenv("KAFKA_PASSWORD"))

   if err !=nil{
	log.Fatalln(err)
   }

   dialer := &kafka.Dialer{
    SASLMechanism: mechanism,
    TLS:           &tls.Config{},
  }

    reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_ENDPOINT")},
		Topic: os.Getenv("KAFKA_TOPIC"),
		Dialer: dialer,
	})
	defer reader.Close()
    
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel()
    
	message, err := reader.ReadMessage(ctx)
	if err != nil {
		log.Fatalln(err)
	  }
	  
	  var commentObj comment
	 err = json.Unmarshal(message.Value,&commentObj)
	 if err !=nil{
		log.Fatalln(err)
	 }
	log.Printf("%+v\n", commentObj)
	sentiments.AnalyzeSentiment(commentObj.Comment)
}