package ral

import (
	"fmt"
	"strconv"
	"strings"
	"encoding/csv"
	"encoding/json"
	"os"

	"github.com/eidolon/wordwrap"
)

// Describes an output format
type Format int
const (
	FormatSimple Format = iota
	FormatCSV
	FormatShArray
	FormatJson )

// Surround string with single quotes
func SurroundQuoteSingle(s string) (string) {
	return strings.Join([]string{"'", s, "'"}, "") }

// Escape single quotes in a string for shell
func EscapeShSingle(s string) (string) {
	return strings.ReplaceAll(s, "'", "'\"'\"'") }

// Single quote a phrase for shell
func ShQuote(s string) (string) {
	return SurroundQuoteSingle(EscapeShSingle(s)) }

// Serialize ContinuityList to console
func (cl ContinuityList) Print(f Format) {
	switch(f) {
	case FormatShArray:
		for _, c := range cl {
			fmt.Printf("(%s %s %d)\n",
				ShQuote(c.Name),
				ShQuote(c.Description),
				c.PostCount)
		}
	case FormatCSV:
		writer := csv.NewWriter(os.Stdout)
		for _, c := range cl {
			writer.Write([]string{
				c.Name,
				c.Description,
				strconv.Itoa(c.PostCount) })
		}
		writer.Flush()
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
	case FormatShArray:
		for _, t := range tl {
			fmt.Printf("(%s %d %d %s %d %s)\n",
				ShQuote(t.Continuity),
				t.Year,
				t.Topic,
				ShQuote(t.Created),
				t.Replies,
				ShQuote(t.Content))
		}
	case FormatCSV:
		writer := csv.NewWriter(os.Stdout)
		for _, t := range tl {
			writer.Write([]string{
				t.Continuity,
				strconv.Itoa(t.Year),
				strconv.Itoa(t.Topic),
				t.Created,
				strconv.Itoa(t.Replies),
				t.Content })
		}
		writer.Flush()
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
	case FormatShArray:
		for _, y := range yl {
			fmt.Printf("(%s %d %d)\n",
				ShQuote(y.Continuity),
				y.Year,
				y.Count)
		}
	case FormatCSV:
		writer := csv.NewWriter(os.Stdout)
		for _, y := range yl {
			writer.Write([]string{
				y.Continuity,
				strconv.Itoa(y.Year),
				strconv.Itoa(y.Count) })
		}
		writer.Flush()
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
