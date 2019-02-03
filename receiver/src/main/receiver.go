package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mkBody(key string, value interface{}) map[string]interface{} {
	return map[string]interface{}{
		"body": map[string]interface{}{
			key: value,
		},
		"headers": map[string]interface{}{
			"Content-Type": "application/json",
		},
	}
}

// generate a random number to be used as grop name
var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

// Main received messages using a consumer for nickname
func Main(args map[string]interface{}) map[string]interface{} {

	var topic string
	var partition float64
	var ok bool

	// check if there are the topic and the partition
	topic, ok = args["topic"].(string)
	if !ok {
		return mkBody("error", "topic required")
	}
	partition, ok = args["partition"].(float64)
	if !ok {
		return mkBody("error", "partition required")
	}

	// handle password protection
	if pass, ok := args["pass"].(string); ok {
		if pass == args["secret"] {
			// authorized, generate group
			group := fmt.Sprintf("g%d", generator.Uint64())
			return mkBody("group", group)
		}
		return mkBody("error", "bad password")
	}

	// check there is the group then handle receive
	if group, ok := args["group"].(string); ok {
		// get the consumer and read the messages
		config := configConsumer(args)
		consumer := Consumer(config, topic, int32(partition), group)
		return mkBody("messages", Receive(consumer))
	}
	return mkBody("error", "group required")
}
