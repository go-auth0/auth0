package management

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/rehttp"
	"golang.org/x/oauth2"
)

func wrapRetry(c *http.Client) *http.Client {
	return &http.Client{
		Transport: rehttp.NewTransport(
			c.Transport,
			func(attempt rehttp.Attempt) bool {
				if attempt.Response == nil {
					return false
				}
				return attempt.Response.StatusCode == http.StatusTooManyRequests
			},
			func(attempt rehttp.Attempt) time.Duration {
				resetAt := attempt.Response.Header.Get("X-RateLimit-Reset")
				resetAtUnix, err := strconv.ParseInt(resetAt, 10, 64)
				if err != nil {
					resetAtUnix = time.Now().Add(5 * time.Second).Unix()
				}
				return time.Unix(resetAtUnix, 0).Sub(time.Now())
			},
		),
	}
}

func wrapUserAgent(c *http.Client) *http.Client {
	if c == nil {
		c = http.DefaultClient
	}
	return &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Set("User-Agent", "Go-Auth0-SDK/v0")
			return c.Transport.RoundTrip(r)
		}),
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (rf roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return rf(r)
}

// authConfig is the payload used to receive an Auth0 management token. This token
// is a JWT, it contains specific granted permissions (known as scopes), and it
// is signed with a application API key and secret for the entire tenant.
//
// 	{
// 	  "audience": "https://YOUR_AUTH0_DOMAIN/api/v2/",
// 	  "client_id": "YOUR_CLIENT_ID",
// 	  "client_secret": "YOUR_CLIENT_SECRET",
// 	  "grant_type": "client_credentials"
// 	}
//
// See: https://auth0.com/docs/api/management/v2/tokens#1-get-a-token
//
type authConfig struct {
	Audience     string `json:"audience"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

// token is the response body from the request to receive an Auth0 management
// token.
//
// 	{
// 	  "access_token": "eyJ...Ggg",
// 	  "expires_in": 86400,
// 	  "scope": "read:clients create:clients read:client_keys",
// 	  "token_type": "Bearer"
// 	}
//
// See: https://auth0.com/docs/api/management/v2/tokens#2-use-the-token
//
type token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// tokenSource is the source
type tokenSource struct {
	domain     string
	authConfig *authConfig
}

func (ts *tokenSource) Token() (*oauth2.Token, error) {

	var payload bytes.Buffer
	json.NewEncoder(&payload).Encode(ts.authConfig)

	req, _ := http.NewRequest("POST", "https://"+ts.domain+"/oauth/token", &payload)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, newError(res.Body)
	}

	t := &token{}

	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(t); err != nil {
		return nil, err
	}

	return &oauth2.Token{
		AccessToken: t.AccessToken,
		TokenType:   t.TokenType,
		Expiry:      time.Now().Add(time.Duration(t.ExpiresIn) * time.Second),
	}, nil
}

func newClient(domain, clientID, clientSecret string) (*http.Client, error) {

	src := &tokenSource{
		domain: domain,
		authConfig: &authConfig{
			Audience:     "https://" + domain + "/api/v2/",
			ClientID:     clientID,
			ClientSecret: clientSecret,
			GrantType:    "client_credentials",
		},
	}

	ts := oauth2.ReuseTokenSource(nil, src)

	return oauth2.NewClient(context.Background(), ts), nil
}
