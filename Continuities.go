package ral

import (
	"encoding/json"
	"fmt"
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
	params := map[string]string{}
	areq := s.APIRequest(View, params)

	body, err := areq.Go()
	if err != nil { return }

	err = json.Unmarshal(body, &ret)
	if err != nil { return }

	return
}

// Serialize to console
func (cl ContinuityList) Print(f Format) {
	switch(f) {
	case FormatSimple:
		for i, c := range cl {
			fmt.Printf("%d. [%s]\n", i+1, c.Name)
			fmt.Printf("    %s\n", c.Description)
			fmt.Printf("    %d posts\n", c.PostCount)
		}
	} }
