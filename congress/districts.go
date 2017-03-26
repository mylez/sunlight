package congress

import (
	"fmt"
	"net/http"
)

type District struct {
	State    string `json:"state"`
	District int    `json:"district"`
}

type DistrictResult struct {
	Result
	Results []District `json:"results"`
}

func (c *Congress) GetDistrictsByLatLon(lat float32, lon float32) (*DistrictResult, error) {
	var params = map[string]string{
		"latitude":  fmt.Sprintf("%f", lat),
		"longitude": fmt.Sprintf("%f", lon),
	}

	l := DistrictResult{}
	err := c.GetURL(&l, congressRoot, params, "districts", "locate")
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func (c *Congress) GetDistrictsByZip(zip string) (*DistrictResult, error) {
	var params = map[string]string{
		"zip":  zip,
	}
	l := DistrictResult{}
	err := c.GetURL(&l, congressRoot, params, "districts", "locate")
	if err != nil {
		return nil, err
	}
	return &l, nil
}


func (c *Congress) GetDistrictsByZip_Client(client *http.Client, zip string) (*DistrictResult, error) {
	var params = map[string]string{
		"zip":  zip,
	}
	l := DistrictResult{}
	err := c.GetURL(&l, congressRoot, params, "districts", "locate")
	if err != nil {
		return nil, err
	}
	return &l, nil
}
