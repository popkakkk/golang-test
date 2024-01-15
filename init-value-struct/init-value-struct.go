package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name"`
}

type Message struct {
	Text string `json:"text"`
}

func main() {
	var resp struct {
		Result struct {
			Name string
		}
		Responsemsge string
		Responsecode string
	}

	resp.Responsecode = "98"

	jsonn := string(`{
		"Result":{
		    "id" : 2,
		   "name":"pop patipan"
		}
	 }`)

	err := json.Unmarshal([]byte(jsonn), &resp)
	if err != nil {
		panic(err)
	}

	fmt.Println("resp", resp)

	var a A
	a.ID = 1
	str := string(`{
	
		"Name" : "Pop"
	 }`)

	err = json.Unmarshal([]byte(str), &a)
	if err != nil {
		panic(err)
	}

	fmt.Println("a", a)

}
