package goharvest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

type Client struct {
	// The base path of the Harvest API, as detailed on Harvest's
	// documentation site
	BasePath string

	// The token to use when making requests, provided by the application.
	// Harvest provides OAuth or PAT style authentication.
	Token string

	// The HTTP Client that does the heavy lifting
	Client http.Client

	// The AccountID against which the client will make requests.
	AccountID string

	// The UserAgent to use when making requests. This is a string that
	// typically takes the form of "<appName> (<contact>)"; i.e.:
	// "MyHarvestApp (https://www.myharvestapp.com/contact)" or
	// "MyHarvestApp (myemail@email.com)"
	UserAgent string
}

// Create a new Client with the provided token, account ID, and
// User-Agent string
func NewClient(PAT string, accountID string, userAgent string) *Client {
	return &Client{
		Token:     PAT,
		Client:    http.Client{},
		BasePath:  "https://api.harvestapp.com",
		AccountID: accountID,
		UserAgent: userAgent,
	}
}

// Make a GET request to the client's BasePath + the provided urlTail
func (c *Client) Get(urlTail string) (*http.Response, error) {
	return c.makeRequest("GET", urlTail, nil)
}

// Make a POST request to the client's BasePath + the provided urlTail,
// using the provided request body.
func (c *Client) Post(urlTail string, body any) (*http.Response, error) {
	return c.makeRequest("POST", urlTail, body)
}

func (c *Client) Patch(urlTail string, body any) (*http.Response, error) {
	return c.makeRequest("PATCH", urlTail, body)
}

func (c *Client) Delete(urlTail string) error {
	_, err := c.makeRequest("DELETE", urlTail, nil)
	return err
}

// Creates a new request with the provided method, urlTail, and body,
// setting the appropriate headers each time.
func (c *Client) newRequest(method string, urlTail string, body any) (*http.Request, error) {
	url := c.BasePath + urlTail
	var req *http.Request
	var err error
	if body == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		var bodyReader bytes.Buffer
		err := json.NewEncoder(&bodyReader).Encode(body)
		if err != nil {
			return req, err
		}
		req, err = http.NewRequest(method, url, &bodyReader)
	}
	if err != nil {
		return req, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("Harvest-Account-Id", c.AccountID)
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Content-Type", "application/json")

	return req, err
}

// A struct to carry error codes from the response up to the
// main application
type ErrorCodeResponse struct {
	StatusCode int
	Message    string
}

func (e ErrorCodeResponse) Error() string {
	return fmt.Sprintf("Error: %d, %s", e.StatusCode, e.Message)
}

// Crate and issue a request with the client. This wrapper also checks for
// any error response codes and creates and returns the appropriate error
// object.
func (c *Client) makeRequest(method string, urlTail string, body any) (*http.Response, error) {
	req, err := c.newRequest(method, urlTail, body)
	if err != nil {
		return &http.Response{}, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return res, err
	}

	// Handle non-200 results
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		// io.ReadAll() will consume the response body, so we need to re-set it
		ba, err := io.ReadAll(res.Body)
		res.Body = io.NopCloser(bytes.NewBuffer(ba))
		if err != nil {
			return res, err
		}
		return res, ErrorCodeResponse{res.StatusCode, string(ba)}
	}

	return res, nil
}

// Add query params, if needed, to a urlTail
func buildPathWithParams[T any](urlTail string, params T) (string, error) {
	qs, err := query.Values(params)
	if err != nil {
		return "", err
	}
	queryString := qs.Encode()
	if queryString != "" {
		return urlTail + "?" + queryString, nil
	}
	return urlTail, nil
}
