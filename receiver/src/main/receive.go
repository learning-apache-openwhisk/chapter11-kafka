package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var consumers = make(map[string]*kafka.Consumer)
var passwords = make(map[string]string)

// Config builds a config map from args
func Config(args map[string]interface{}) *kafka.ConfigMap {

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

// Consumer return a consumer by nick and check the password
func Consumer(config *kafka.ConfigMap, topic string, partition int32, nick string, pass string) *kafka.Consumer {

	// return cached consumer, if any
	if consumer, ok := consumers[nick]; ok {
		if passwords[nick] == pass {
			return consumer
		}
		return nil
	}

	// not found in cache,
	// create a consumer and return it
	config.SetKey("group.id", nick)
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Println(err)
		return nil
	}

	// cache values and subscribe
	consumers[nick] = consumer
	passwords[nick] = pass
	// assign to a specific topic and partition
	assignment := []kafka.TopicPartition{{Topic: &topic, Partition: partition}}
	consumer.Assign(assignment)
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
			case kafka.PartitionEOF:
				return messages
			}
		default:
			return messages
		}
	}
}
