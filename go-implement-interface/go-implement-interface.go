package main

import "context"

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

// Nop is an endpoint that does nothing and returns a nil error.
// Useful for tests.
func Nop(context.Context, interface{}) (interface{}, error) { return struct{}{}, nil }

// Middleware is a chainable behavior modifier for endpoints.
type Middleware func(Endpoint) Endpoint

func main() {

}

func giveMiddleware() Middleware {
	return func(Endpoint) Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) { return nil, nil }
	}
}
