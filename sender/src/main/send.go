package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

// Producer returns a producer to Kafka in a persistent way
func Producer(args map[string]interface{}) *kafka.Producer {
	if producer != nil {
		return producer
	}

	// extract broker list from map
	brokers := ""
	for i, s := range args["kafka_brokers_sasl"].([]interface{}) {
		if i == 0 {
			brokers = s.(string)
		} else {
			brokers += "," + s.(string)
		}
	}

	// generate configuration
	config := kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"security.protocol": "sasl_ssl",
		"sasl.mechanisms":   "PLAIN",
		"sasl.username":     args["user"],
		"sasl.password":     args["password"],
	}

	// create a producer and return it
	p, err := kafka.NewProducer(&config)
	if err != nil {
		log.Println(err)
		return nil
	}
	producer = p
	deliveryChan = make(chan kafka.Event)
	return producer
}

var deliveryChan chan kafka.Event

// Send a message
func Send(p *kafka.Producer, topic string, partition int, message []byte) error {
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic, Partition: int32(partition)},
		Value: message,
	}, deliveryChan)
	e := <-deliveryChan
	m := e.(*kafka.Message)
	return m.TopicPartition.Error
}
