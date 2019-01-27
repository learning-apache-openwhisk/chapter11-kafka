package klient

import (
	"fmt"
)

func ExampleSend() {
	args := readCred()
	//fmt.Printf("%v", args)
	p := Producer(args)
	fmt.Printf("%v\n", p)
	err := Send(p, "queue", []byte("hello"), int32(0))
	fmt.Println(err)
	p.Flush(1000)
	// Output:
	// rdkafka#producer-1
	// <nil>
}

/*
func ExampleMain() {
	args := readCred()
	args["topic"] = "queue"
	args["message"] = "hello world"
	res := Main(args)
	fmt.Printf("%v\n", res)
	// Output:
	// map[ok:true]
}
*/
