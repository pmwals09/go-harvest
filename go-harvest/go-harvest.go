package goharvest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	BasePath  string
	Token     string
	Client    http.Client
	AccountID string
	UserAgent string
}

func NewClient(PAT string, accountID string, userAgent string) *Client {
	return &Client{
		Token:     PAT,
		Client:    http.Client{},
		BasePath:  "https://api.harvestapp.com",
		AccountID: accountID,
		UserAgent: userAgent,
	}
}

func (c *Client) GetMyProjectAssignments() (ProjectAssignmentResponse, error) {
	pa := ProjectAssignmentResponse{}
	urlTail := "/v2/users/me/project_assignments"
	res, err := c.Get(urlTail)
	if err != nil {
		return pa, err
	}
	err = json.NewDecoder(res.Body).Decode(&pa)
	if err != nil {
		return pa, err
	}
	return pa, nil
}

func (c *Client) GetMe() (User, error) {
	u := User{}
	urlTail := "/v2/users/me"
	res, err := c.Get(urlTail)
	if err != nil {
		return u, err
	}
	err = json.NewDecoder(res.Body).Decode(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (c *Client) GetTimeEntries() (TimeEntryResponse, error) {
  tr := TimeEntryResponse{}
  res, err := c.Get("/v2/time_entries")
  if err != nil {
    return tr, err
  }
  err = json.NewDecoder(res.Body).Decode(&tr)
  if err != nil {
    return tr, err
  }
  return tr, nil
}

func (c *Client) Get(urlTail string) (*http.Response, error) {
  return c.makeRequest("GET", urlTail, nil)
}

func (c *Client) newRequest(method string, urlTail string, body io.Reader) (*http.Request, error) {
	url := c.BasePath + urlTail
	var req *http.Request
	var err error
	if body == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, body)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("Harvest-Account-Id", c.AccountID)
	req.Header.Set("User-Agent", c.UserAgent)

	return req, err
}

type ErrorCodeResponse struct {
  StatusCode int
  Message string
}
func (e ErrorCodeResponse) Error() string {
  return fmt.Sprintf("Error: %d, %s", e.StatusCode, e.Message)
}
func (c *Client) makeRequest(method string, urlTail string, body io.Reader) (*http.Response, error) {
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

