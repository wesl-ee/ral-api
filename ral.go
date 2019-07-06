// Provides access to the posts and other information on the RAL textboard
// (https://ralee.org) using its API
package ral

import (
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
)

// Defines the user agent projected by the package
const USER_AGENT = "GoRAL"

// Functions exposed by the RAL API
type APIAction string
const (
	View APIAction = "view"
)

// Defines site-wide parameters for accessing information from RAL
type Site struct {
	Endpoint string
	Timeout int
	UserAgent string
}

// Defines a single request to the RAL API
type APIRequest struct {
	Action APIAction
	Endpoint string
	UserAgent string
	Parameters map[string]string
	Client http.Client
}

// Creates a new site with sane defaults
func New() (Site) {
	return Site{
		Timeout: 5,
		UserAgent: USER_AGENT}
}

// Serialize the API request into a URI
func (a APIRequest) URI() (string) {
	var sb strings.Builder
	sb.WriteString(a.Endpoint)
	sb.WriteString("?a=")
	sb.WriteString(string(a.Action))
	for key, val := range a.Parameters {
		sb.WriteString(key)
		sb.WriteString("=")
		sb.WriteString(val)
	}
	fmt.Printf(sb.String())
	return sb.String()
}

// Execute a single request to the RAL API
func (areq APIRequest) Go() (body []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, areq.URI(), nil)
	if err != nil { return }

	req.Header.Set("User-Agent", areq.UserAgent)

	res, err := areq.Client.Do(req)
	if err != nil { return }

	body, err = ioutil.ReadAll(res.Body)
	return
}
