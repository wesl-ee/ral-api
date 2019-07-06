package ral

import (
	"net/http"
	"encoding/json"
	"time"
)

// Continuities represent discussion topics like "Anime" or "Music"
type ContinuityList []Continuity
type Continuity struct {
	Name string
	PostCount int
	Description string
}

// Generate and execute an API request which returns the full list of
// Continuities on the given Site
func (s Site) Continuities() (ret ContinuityList, err error) {
	areq := APIRequest{
		Action: View,
		UserAgent: s.UserAgent,
		Endpoint: s.Endpoint,
		Client: http.Client{
			Timeout: (time.Second * time.Duration(s.Timeout))} }

	body, err := areq.Go()
	if err != nil { return }

	err = json.Unmarshal(body, &ret)
	if err != nil { return }

	return
}
