package main

func mkErr(message string) map[string]interface{} {
	return map[string]interface{}{
		"body": map[string]interface{}{
			"error": message,
		},
		"headers": map[string]interface{}{
			"Content-Type": "application/json",
		},
	}
}

// Main received messages using a consumer for nickname
func Main(args map[string]interface{}) map[string]interface{} {

	var nick string
	var pass string
	var topic string
	var partition float64
	var ok bool

	// check there is the nickname
	if nick, ok = args["nick"].(string); !ok {
		return mkErr("nick required")
	}

	// check there is the nickname
	if pass, ok = args["pass"].(string); !ok {
		return mkErr("pass required")
	}

	// check if there is the topic and the partition
	topic, ok = args["topic"].(string)
	if !ok {
		return mkErr("topic required")
	}
	partition, ok = args["partition"].(float64)
	if !ok {
		return mkErr("partition required")
	}

	// get the consumer and read the messages
	config := Config(args)
	consumer := Consumer(config, topic, int32(partition), nick, pass)
	if consumer == nil {
		return mkErr("wrong password")
	}

	// returning the result
	return map[string]interface{}{
		"body": map[string]interface{}{
			"messages": Receive(consumer),
		},
		"headers": map[string]interface{}{
			"Content-Type": "application/json",
		},
	}
}
