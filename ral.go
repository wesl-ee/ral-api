// Provides access to the posts and other information on the RAL textboard
// (https://ralee.org) using its API
package ral

import (
	"net/http"
	"strings"
	"fmt"
	"net/url"
	"io/ioutil"
	"time"
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
	URL url.URL
	Timeout int
	UserAgent string
}

// Defines a single request to the RAL API
type APIRequest struct {
	URL url.URL
	// Action APIAction
	// Endpoint string
	UserAgent string
	// Parameters map[string]string
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
	return a.URL.String()
}

// Generates a single request for information on the given site, representing
// exactly one transaction / HTTP request
func (s Site) APIRequest(action APIAction, params map[string]string) (a APIRequest) {
	var querySlice []string
	for key, val := range params {
		querySlice = append(querySlice,
			strings.Join([]string{key, val}, "="))
	}
	querySlice = append(querySlice,
		strings.Join([]string{"a", string(action)}, "="))
	rawQuery := strings.Join(querySlice, "&")

	a.URL = url.URL{
		Scheme: s.URL.Scheme,
		User: s.URL.User,
		Host: s.URL.Host,
		Path: s.URL.Path,
		RawQuery: rawQuery}
	a.UserAgent = s.UserAgent
	a.Client = http.Client{
		Timeout: (time.Second * time.Duration(s.Timeout))}
	return
}

// Some day... do the heavy lifting done in APIRequest.URI()
// func CreateQueryString(params map[string]string) string { }

// Execute a single request to the RAL API
func (areq APIRequest) Go() (body []byte, err error) {
	fmt.Println(areq.URI())
	req, err := http.NewRequest(http.MethodGet, areq.URI(), nil)
	if err != nil { return }

	req.Header.Set("User-Agent", areq.UserAgent)

	res, err := areq.Client.Do(req)
	if err != nil { return }

	body, err = ioutil.ReadAll(res.Body)
	return
}
