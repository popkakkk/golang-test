package builtinlib

import "fmt"

type Allower interface {
	Allow() bool
}

func AllowProcess(a Allower) {
	fmt.Println("process", !a.Allow())
}
