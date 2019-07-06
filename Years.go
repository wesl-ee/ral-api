package ral

import (
	"net/http"
	"encoding/json"
	"time"
)

// Years are subsets of Continuities; they represent all Topics posted
// in that Continuity in a given year
type YearList []Year
type Year struct {
	Year int
	Continuity string
	Count int
}

// Generate and execute an API request which returns the full list of
// Years on the given Site and given Continuity
func (s Site) Years(continuity string) (ret YearList, err error) {
	areq := APIRequest{
		Action: View,
		Endpoint: s.Endpoint,
		UserAgent: s.UserAgent,
		Parameters: map[string]string {
			"continuity": continuity},
		Client: http.Client{
			Timeout: (time.Second * time.Duration(s.Timeout))} }

	body, err := areq.Go()
	if err != nil { return }

	err = json.Unmarshal(body, &ret)
	return
}