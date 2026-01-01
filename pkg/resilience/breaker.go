package resilience

import (
	"time"

	"github.com/sony/gobreaker"
)

// create a global map for reuse if needed
var Breakers = make(map[string]*gobreaker.CircuitBreaker)

// NewBreaker creates a new circuit breaker with professional settings
func NewBreaker(name string) *gobreaker.CircuitBreaker {
	if b, exists := Breakers[name]; exists {
		return b
	}

	b := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        name,
		MaxRequests: 2,                //allowed requests in half open
		Interval:    60 * time.Second, //reset failure count
		Timeout:     5 * time.Second,  //open state duration
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
	})

	Breakers[name] = b
	return b
}
