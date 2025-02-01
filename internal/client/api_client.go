package client

import (
	"circuit-break/internal/utils"
	"fmt"
	"log"
	"time"

	"github.com/sony/gobreaker"
)

type ApiClient struct {
	CircuitBreaker *gobreaker.CircuitBreaker
}

func NewApiClient() *ApiClient {
	settings := gobreaker.Settings{
		Name:        "HTTP GET",
		MaxRequests: 1,
		Interval:    0,
		Timeout:     1 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 3
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("Circuit breaker state changed from %s to %s\n", from.String(), to.String())
		},
	}
	cb := gobreaker.NewCircuitBreaker(settings)
	return &ApiClient{
		CircuitBreaker: cb,
	}
}

func (c *ApiClient) MakeRequest(url string) (string, error) {
	log.Println("Making request to", url)

	body, err := c.CircuitBreaker.Execute(func() (interface{}, error) {
		response, err := utils.HttpGet(url)
		if err != nil {
			log.Printf("Request failed: %v\n", err)
			return nil, err
		}
		defer response.Body.Close()
		body := utils.ReadResponseBody(response)
		if response.StatusCode == 200 {
			log.Println("Request successful")
			return body, nil
		} else {
			return "", fmt.Errorf("falha servidor intrerno - [%d]", response.StatusCode)
		}
	})
	switch c.CircuitBreaker.State() {
	case gobreaker.StateOpen:
		{
			return "", fmt.Errorf("circuit breaker is open")
		}
	}
	if err != nil {
		log.Printf("Circuit breaker triggered: %v\n", err)
		return "", err
	}

	return body.(string), nil
}
