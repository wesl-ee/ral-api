package ral

import (
	"encoding/json"
	"fmt"
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
	params := map[string]string{
		"continuity": continuity}
	areq := s.APIRequest(View, params)

	body, err := areq.Go()
	if err != nil { return }

	err = json.Unmarshal(body, &ret)
	return
}

// Serialize to console
func (yl YearList) Print(f Format) {
	switch(f) {
	case FormatSimple:
		for i, y := range yl {
			fmt.Printf("%d. [%s/%d]\n", i+1, y.Continuity, y.Year)
			fmt.Printf("    %d posts\n", y.Count)
		}
	} }
