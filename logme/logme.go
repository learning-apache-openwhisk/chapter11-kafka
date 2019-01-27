package main

import "log"
import "encoding/json"

// Main is a logger function
func Main(args map[string]interface{}) map[string]interface{} {
	obj, _ := json.Marshal(args)
	log.Printf("%s\n", obj)
	return args
}
