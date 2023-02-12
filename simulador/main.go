package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafka2 "github.com/eduardotecnologo/imersao.fullcycle23/application/kafka"
	"github.com/eduardotecnologo/imersao.fullcycle23/infra/kafka"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
	//producer := kafka.NewKafkaProducer()
	// kafka.Publish("Ola", "readteste", producer)

	// for {
	// 	_ = 1
	// }
	//	route := route.Route{            02:07:00
	// 	ID:       "1",
	// 	ClientID: "1",
	// }
	// route.LoadPositions()
	// stringjson, _ := route.ExportJsonPositions()
	// fmt.Println(stringjson[0])
}
