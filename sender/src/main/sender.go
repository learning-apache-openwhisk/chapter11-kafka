package main

func mkErr(message string) map[string]interface{} {
	return map[string]interface{}{
		"body": "ERROR: " + message,
	}
}

// Main is the entry point
func Main(args map[string]interface{}) map[string]interface{} {

	// get args
	message, ok := args["message"].(string)
	if !ok {
		return mkErr("no message")
	}
	topic, ok := args["topic"].(string)
	if !ok {
		return mkErr("no topic")
	}
	partition, ok := args["partition"].(float64)
	if !ok {
		return mkErr("no partition")
	}

	// retrieving the connection
	p := Producer(args)
	if p == nil {
		return mkErr("cannot connect")
	}

	// sending the message
	//part, _ := strconv.Atoi(partition)
	err := Send(p, topic, int(partition), []byte(message))
	if err != nil {
		return mkErr(err.Error())
	}
	p.Flush(10)

	return map[string]interface{}{
		"body": "OK",
	}
}
