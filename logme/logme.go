package main

import "log"
import "encoding/json"

// Main is a logger function
func Main(args map[string]interface{}) map[string]interface{} {
	jo, _ := json.Marshal(args)
	log.Printf("%s\n", jo)
	return args
}
