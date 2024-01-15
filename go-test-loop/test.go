package main

import "fmt"

type AA struct {
	A A
}
type A struct {
	id             int
	progress_value int
}

func main() {

	as := []A{{id: 1, progress_value: 0}, {id: 2, progress_value: 2}}

	var aas []AA

	for _, a := range as {

		if !cal(&a) {
			continue
		}

		createA := AA{}
		createA.A = a

		aas = append(aas, createA)
	}

	fmt.Println("aas", aas)
}

func cal(a *A) bool {

	if a.id == 2 {
		return true
	}

	if a.id == 1 {
		a.progress_value = 0
	}

	return false
}
