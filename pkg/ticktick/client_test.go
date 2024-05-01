package ticktick_test

import (
	"testing"

	"github.com/mastorm/ynab-to-ticktick/pkg/ticktick"
)

func TestAuthorizeUrl(t *testing.T) {
	const (
		clientId    = "foo"
		accessToken = "baz"
		state       = "stateFoo"
		redirectUri = "https://foo-bar.com"
	)

	scopes := []ticktick.Scope{ticktick.TasksRead, ticktick.TasksWrite}

	args := ticktick.ClientArgs{ClientId: clientId, AccessToken: accessToken, Scopes: scopes}
	client := ticktick.NewClient(args)
	url, err := client.AuthorizeUrl(state, redirectUri)
	if err != nil {
		t.Error(err)
	}

	if url.Host != "ticktick.com" {
		t.Errorf("url.host is %s, want: ticktick.com", url.Host)
	}

	q := url.Query()
	probeKey := func(key string, expected string) {
		value := q.Get(key)
		if value != expected {
			t.Errorf("Expected %s to be %s, got: %s", key, expected, value)
		}
	}

	probeKey("redirect_uri", redirectUri)
	probeKey("state", state)
	probeKey("client_id", clientId)
	probeKey("scope", "tasks:read tasks:write")
}
