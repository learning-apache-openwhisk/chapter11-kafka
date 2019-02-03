package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func configProducer(args map[string]interface{}) *kafka.ConfigMap {

	// extract broker list from map
	brokers := ""
	for _, s := range args["kafka_brokers_sasl"].([]interface{}) {
		brokers += s.(string) + ","
	}
	brokers = brokers[0 : len(brokers)-1]

	// generate configuration
	config := kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"security.protocol": "sasl_ssl",
		"sasl.mechanisms":   "PLAIN",
		"sasl.username":     args["user"],
		"sasl.password":     args["password"],
	}
	return &config
}

var producer *kafka.Producer
var deliveryChan chan kafka.Event

// Producer returns a producer to Kafka in a persistent way
func Producer(args map[string]interface{}) *kafka.Producer {
	if producer != nil {
		return producer
	}

	// create a producer and return it
	p, err := kafka.NewProducer(configProducer(args))
	if err != nil {
		log.Println(err)
		return nil
	}
	producer = p
	deliveryChan = make(chan kafka.Event)
	return producer
}

// Send a message
func Send(p *kafka.Producer, topic string, partition int, message []byte) error {
	tp := kafka.TopicPartition{
		Topic:     &topic,
		Partition: int32(partition),
	}
	msg := &kafka.Message{
		TopicPartition: tp,
		Value:          message,
	}
	p.Produce(msg, deliveryChan)
	e := <-deliveryChan
	m := e.(*kafka.Message)
	return m.TopicPartition.Error
}
