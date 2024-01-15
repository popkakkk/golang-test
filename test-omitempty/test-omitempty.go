package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	ID    int64
	Class string `json:"class,omitempty"`
}

type UserService interface {
	GetUser(string) User
	SaveUser()
}

type ServiceA struct {
	Code string
}

func newServiceA() ServiceA {
	return ServiceA{}

}

func (s ServiceA) GetUser(string) User {
	return User{Name: "Service A name"}
}

func (s ServiceA) SaveUser() {

}

type ServiceB struct {
	ContractID string
}

func newServiceB() *ServiceB {
	return &ServiceB{}
}

func (s *ServiceB) GetUser(a string) User {
	return User{Name: a}
}

func (s ServiceB) SaveUser() {

}

func main() {
	fmt.Println("start")

	userA := User{
		ID:   1,
		Name: "Pop",
	}
	fmt.Println("init userA", userA)

	b, err := json.Marshal(userA)
	if err != nil {
		panic(err)
	}

	fmt.Println("json userA", string(b))
	fmt.Println("----------------")

	aService := newServiceA()
	display(aService)
	fmt.Println("stop")
	fmt.Println("---------------")

	bService := newServiceB()
	display(bService)

}

func display(s UserService) {
	host := "localhost"
	fmt.Println("GetUser println", s.GetUser(host))
}
