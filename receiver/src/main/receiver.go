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

var passwords = map[string]string{}
var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

// Login checks for authentication and return a group for each user
// cache the password of a never-seen-before user for reuse
func Login(nick string, pass string) map[string]interface{} {
	if password, ok := passwords[nick]; ok {
		if pass != password {
			return mkBody("error", "bad username or password")
		}
	} else {
		passwords[nick] = pass
	}
	group := fmt.Sprintf("g%d", generator.Uint64())
	return mkBody("group", group)
}

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

	// handle nickname protection
	if nick, ok := args["nick"].(string); ok {
		if pass, ok := args["pass"].(string); ok {
			return Login(nick, pass)
		}
		return mkBody("error", "nick and pass required")
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
