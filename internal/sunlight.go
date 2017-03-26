package sunlight

import (
	"net/http"
	"strings"
	"net/url"
	"encoding/json"
	"io/ioutil"
)

type Sunlight struct {
	apiKey string
	client *http.Client
}

func (s *Sunlight) SetAPIKey(key string){
	s.apiKey = key
}

func (s *Sunlight) SetClient(client *http.Client) {
	s.client = client
}


func (s *Sunlight) GenerateURL(root string, resource ...string) string {
	return root + "/" + strings.Join(resource, "/")
}

func (s *Sunlight) QueryURL(root string, params map[string]string, resource ...string) string {
	uri := s.GenerateURL(root, resource...) + "?apikey=" + s.apiKey
	for k, v := range params {
		uri = uri + "&" + url.QueryEscape(k) + "=" + url.QueryEscape(v)
	}
	return uri
}

func (s *Sunlight) GetURL(target interface{}, root string, params map[string]string, resource ...string) (err error) {
	url := s.QueryURL(root, params, resource...)
	res, err := s.client.Get(url)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, target)
	return
}
