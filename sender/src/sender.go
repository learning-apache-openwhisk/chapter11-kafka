package main

import (
	"encoding/json"
	"fmt"
	"klient"
)

// mkErr makes an error from different sources
func mkErr(message interface{}) map[string]interface{} {
	result := "unknown error"
	switch v := message.(type) {
	case string:
		result = v
	case error:
		result = v.Error()
	}
	return map[string]interface{}{
		"body": "ERROR: " + result,
	}
}

// Main is the entry point
func Main(args map[string]interface{}) map[string]interface{} {
	// retrieving the connection
	p := klient.Producer(args)
	if p == nil {
		return mkErr("cannot connect")
	}
	// getting the topic
	t, ok := args["topic"].(string)
	if !ok {
		return mkErr("no topic")
	}
	m, ok := args["message"]
	if !ok {
		return mkErr("no message")
	}
	// getting the message
	msg, err := json.Marshal(m)
	if err != nil {
		return mkErr(err)
	}
	// sending the message
	fmt.Printf("sending %s -> %s", msg, t)
	err = klient.Send(p, t, msg, 0)
	if err != nil {
		return mkErr(err)
	}
	return map[string]interface{}{
		"body": "OK",
	}
}
