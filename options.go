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
