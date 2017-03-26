package congress

import (
	"github.com/mylez/sunlight/internal"
	"net/http"
)

var congressRoot string = "http://congress.api.sunlightfoundation.com"

type Congress struct {
	sunlight.Sunlight
}

func NewCongress(client *http.Client, apiKey string) *Congress {
	c := &Congress{}

	c.SetAPIKey(apiKey)
	c.SetClient(client)

	return c
}

type Result struct {
	Count int `json:"count"`
	Page  struct {
		  Count   int `json:"count"`
		  PerPage int `json:"per_page"`
		  Page    int `json:"page"`
	} `json:"page"`
}
