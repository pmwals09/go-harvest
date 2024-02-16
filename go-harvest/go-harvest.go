package goharvest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Client struct {
	BasePath string
	Token    string
	Client   http.Client
  AccountID string
  Email string
}

func NewClient(PAT string, accountID string, email string) *Client {
	return &Client{
		Token:  PAT,
		Client: http.Client{},
    BasePath: "https://api.harvestapp.com",
    AccountID: accountID,
    Email: email,
	}
}

func (c *Client) GetMe() (User, error) {
  u := User{}
  urlTail := "/v2/users/me"
  req, err := c.newRequest("GET", urlTail, nil)
  if err != nil {
    return u, err
  }
  res, err := c.Client.Do(req)
  err = json.NewDecoder(res.Body).Decode(&u)
  if err != nil {
    return u, err
  }
  return u, nil
}

func printError(msgs ...string) {
	fmt.Fprint(os.Stderr, msgs)
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
  req.Header.Set("User-Agent", "GoClient (" + c.Email + ")")

  return req, err 
}
