package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
)

type User struct {
	Name   string
	Price  float64
	Price2 float32
}

func main() {

	num, _ := strconv.ParseFloat("1928.16", 64)
	num2, _ := strconv.ParseFloat("1928.16", 32)
	fmt.Println(num)
	fmt.Println(num2)
	user := User{Name: "Frank", Price: num2, Price2: float32(num2)}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Printf("---> %s \n", string(b))

	c := float64(num2)
	fmt.Println("c float32 to float64 --> \n", c)

	// a := float32(1928.16)
	// b := float64(a)

	// fmt.Println("a", a)
	// fmt.Println("b", b)

}

func EncodeB64(message string) (retour string) {
	base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(base64Text, []byte(message))
	return string(base64Text)
}
