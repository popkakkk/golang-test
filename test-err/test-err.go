package main

import "fmt"

func main() {
	var err error

	if err == nil {
		fmt.Println("err", err)
	}

	err = fmt.Errorf("gg")
	if err != nil {
		fmt.Println("gg", err)
	}

}
