package main

import "fmt"

type User struct {
	Id   int32
	Name string
}

func main() {
	user := User{
		Id:   1,
		Name: "pop",
	}
	fmt.Println("user", user)
}
