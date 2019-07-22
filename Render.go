package ral

import (
	"fmt"
	"encoding/json"

	"github.com/eidolon/wordwrap"
)

// Describes an output format
type Format int
const (
	FormatSimple Format = iota
	FormatJson )

// Serialize ContinuityList to console
func (cl ContinuityList) Print(f Format) {
	switch(f) {
	case FormatSimple:
		for i, c := range cl {
			fmt.Printf("%d. [%s]\n", i+1, c.Name)
			fmt.Printf("    %s\n", c.Description)
			fmt.Printf("    %d posts\n", c.PostCount)
		}
	case FormatJson:
		c, err := json.Marshal(cl)
		if err != nil { panic(err) }
		fmt.Println(string(c))
	} }

// Serialize TopicList to console
func (tl TopicList) Print(f Format) {
	switch(f) {
	case FormatSimple:
		for _, t := range tl {
			fmt.Printf("[%s/%d/%d] (%s) (%d replies)\n",
				t.Continuity,
				t.Year,
				t.Topic,
				t.Created,
				t.Replies)

			wrapper := wordwrap.Wrapper(76, false)
			fmt.Printf("%s\n", wordwrap.Indent(wrapper(t.Content), "    ", true))
		}
	case FormatJson:
		t, err := json.Marshal(tl)
		if err != nil { panic(err) }
		fmt.Println(string(t))
	} }

// Serialize YearList to console
func (yl YearList) Print(f Format) {
	switch(f) {
	case FormatSimple:
		for i, y := range yl {
			fmt.Printf("%d. [%s/%d]\n", i+1, y.Continuity, y.Year)
			fmt.Printf("    %d posts\n", y.Count)
		}
	case FormatJson:
		y, err := json.Marshal(yl)
		if err != nil { panic(err) }
		fmt.Println(string(y))
	} }
