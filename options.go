package axcelerate

import (
	"net/http"
	"time"
)

type Option func(s *Settings)

// // PublisherOptions are used to describe a publisher's configuration.
// // Logger is a custom logging interface.
// type Option struct {
// 	BaseURL    string
// 	Rate       Logger
// 	HttpClient *http.Client
// }

func BaseURL(baseURL string) Option {
	return func(s *Settings) {
		s.baseURL = baseURL
	}
}

func HttpClient(httpClient *http.Client) Option {
	return func(s *Settings) {
		s.httpClient = httpClient
	}
}

func RateLimit(rate int) Option {
	return func(s *Settings) {
		s.rate = rate
	}
}
func RatePer(ratePer time.Duration) Option {
	return func(s *Settings) {
		s.ratePer = ratePer
	}
}

type Settings struct {
	baseURL    string
	httpClient *http.Client
	rate       int
	ratePer    time.Duration
}

// Add this to your options.go file (or wherever your Option funcs live)

// Timeout sets the timeout for the HTTP client. If HttpClient is provided, it modifies that client. Otherwise, it's used when creating the default client.
func Timeout(timeout time.Duration) Option {
	return func(s *Settings) {
		if s.httpClient != nil {
			s.httpClient.Timeout = timeout
		} else {
			// Store a dummy client with the timeout, will be used in NewClient
			s.httpClient = &http.Client{Timeout: timeout}
		}
	}
}

// Example usage:
// c, _ := NewClient(token, ws, Timeout(10*time.Second))
// or, override other options as needed
// c, _ := NewClient(token, ws, HttpClient(myClient), Timeout(25*time.Second))
