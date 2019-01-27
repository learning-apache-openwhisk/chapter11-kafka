package klient

import (
	"encoding/json"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readCred() map[string]interface{} {
	cred, err := ioutil.ReadFile("../../../cred.json")
	check(err)
	res := make(map[string]interface{})
	err = json.Unmarshal(cred, &res)
	check(err)
	return res
}
