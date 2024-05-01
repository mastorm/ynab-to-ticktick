package ticktick

import (
	"net/url"
	"strings"
)

type Client struct {
	ClientArgs
	bearerToken *string
}

type ClientArgs struct {
	AccessToken string
	ClientId    string
	Scopes      []Scope
}

type Scope string

const (
	TasksRead  Scope = "tasks:read"
	TasksWrite Scope = "tasks:write"
)

func NewClient(args ClientArgs) *Client {
	return &Client{
		ClientArgs:  args,
		bearerToken: nil,
	}
}

const baseUrl = "https://ticktick.com"

func (c *Client) stringifyScopes() string {
	scopes := []string{}
	for _, v := range c.Scopes {
		scopes = append(scopes, string(v))
	}

	return strings.Join(scopes, " ")
}

func (c *Client) AuthorizeUrl(state string, redirectUri string) (*url.URL, error) {
	url, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	url.Path = "/oauth/authorize"

	q := url.Query()
	q.Add("scope", c.stringifyScopes())
	q.Add("client_id", c.ClientArgs.ClientId)
	q.Add("state", state)
	q.Add("redirect_uri", redirectUri)
	q.Add("response_type", "code")

	url.RawQuery = q.Encode()
	return url, nil
}
