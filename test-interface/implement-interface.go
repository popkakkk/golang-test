package main

import (
	"github.com/popkakkk/test-interface/builtinlib"
	"github.com/popkakkk/test-interface/liba"
)

func main() {
	allow := true
	builtinlib.AllowProcess(liba.NewLibA(allow))
}
