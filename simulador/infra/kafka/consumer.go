package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaCosumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

// Método de consumo, para ler um tópico do kafka
func (k *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error creating consumer kafka message" + err.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopic")}
	c.SubscribeTopics(topics, nil)

	fmt.Println("Kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}
