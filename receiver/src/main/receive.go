package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func configConsumer(args map[string]interface{}) *kafka.ConfigMap {

	// extract broker list from map
	brokers := ""
	for _, s := range args["kafka_brokers_sasl"].([]interface{}) {
		brokers += s.(string) + ","
	}
	brokers = brokers[0 : len(brokers)-1]

	// generate configuration
	config := kafka.ConfigMap{
		"bootstrap.servers":               brokers,
		"security.protocol":               "sasl_ssl",
		"sasl.mechanisms":                 "PLAIN",
		"sasl.username":                   args["user"],
		"sasl.password":                   args["password"],
		"auto.offset.reset":               "latest",
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": false,
		"enable.partition.eof":            true,
		"enable.auto.commit":              true,
	}
	return &config
}

var consumers = map[string]*kafka.Consumer{}

// Consumer return a consumer by nick and check the password
func Consumer(config *kafka.ConfigMap, topic string, partition int32, group string) *kafka.Consumer {

	// return cached consumer, if any
	if consumer, ok := consumers[group]; ok {
		log.Printf("retrieved consumer %s", group)
		return consumer
	}

	// not found in cache,
	// create a consumer and return it
	config.SetKey("group.id", group)
	log.Printf("config %v", config)
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Println(err)
		return nil
	}

	// assign to a specific topic and partition
	assignment := []kafka.TopicPartition{{
		Topic:     &topic,
		Partition: partition}}
	consumer.Assign(assignment)
	consumers[group] = consumer

	// cache values and subscribe
	log.Printf("created consumer %s for %s:%d ", group, topic, partition)
	return consumer
}

// Receive messages
func Receive(c *kafka.Consumer) []string {
	messages := []string{}
	for {
		select {
		case ev := <-c.Events():
			switch e := ev.(type) {
			case *kafka.Message:
				messages = append(messages, string(e.Value))
			}
		default:
			return messages
		}
	}
}
