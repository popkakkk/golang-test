package main

// "net/http"
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/time/rate"
)

func main() {

	service := newHelloService()

	var pingEndpoint endpoint.Endpoint
	pingEndpoint = makePingEndpoint(service)
	pingEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(20000, 20000))(pingEndpoint)
	pingEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(1*rate.Every(time.Second), 1))(pingEndpoint)

	pingHandler := httptransport.NewServer(pingEndpoint, httptransport.NopRequestDecoder, httptransport.EncodeJSONResponse)

	http.Handle("/ping", pingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func makePingEndpoint(s HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fmt.Println("tick")
		return map[string]string{"response": s.ping()}, nil
	}
}

type HelloService interface {
	ping() string
}

type HelloServiceImpl struct {
}

func (s HelloServiceImpl) ping() string {
	time.Sleep(time.Second * 5)

	return "pong"
}

func newHelloService() HelloService {
	return HelloServiceImpl{}
}
