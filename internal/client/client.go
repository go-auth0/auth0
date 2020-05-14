package client

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"

	"github.com/PuerkitoBio/rehttp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"gopkg.in/auth0.v4"
)

// UserAgent is the default user agent string
var UserAgent = fmt.Sprintf("Go-Auth0-SDK/%s", auth0.Version)

func WrapRateLimit(c *http.Client) *http.Client {
	c.Transport = rehttp.NewTransport(
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
			return time.Duration(resetAtUnix-time.Now().Unix()) * time.Second
		},
	)
	return c
}

type Transport struct {
	base    http.RoundTripper
	headers map[string]string
	debug   bool
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range t.headers {
		req.Header.Set(key, value)
	}
	if t.debug {
		dumpRequest(req)
		res, err := t.base.RoundTrip(req)
		if err != nil {
			return res, err
		}
		dumpResponse(res)
		return res, nil
	}
	return t.base.RoundTrip(req)
}

func NewTransport(base http.RoundTripper, headers map[string]string, debug bool) *Transport {
	if base == nil {
		base = http.DefaultTransport
	}
	return &Transport{base, headers, debug}
}

func WrapUserAgent(c *http.Client, userAgent string) *http.Client {
	c.Transport = NewTransport(c.Transport, map[string]string{"User-Agent": userAgent}, false)
	return c
}

func dumpRequest(r *http.Request) {
	b, _ := httputil.DumpRequestOut(r, true)
	log.Printf("\n%s\n", b)
}

func dumpResponse(r *http.Response) {
	b, _ := httputil.DumpResponse(r, true)
	log.Printf("\n%s\n\n", b)
}

func WrapDebug(c *http.Client, debug bool) *http.Client {
	if !debug {
		return c
	}
	c.Transport = NewTransport(c.Transport, map[string]string{}, debug)
	return c
}

func WrapHTTPClientWithOauth2(ctx context.Context, httpClient *http.Client, clientCredentials *clientcredentials.Config) *http.Client {
	ctx = context.WithValue(oauth2.NoContext, oauth2.HTTPClient, httpClient)
	client := oauth2.NewClient(ctx, clientCredentials.TokenSource(ctx))
	httpClient.Transport = client.Transport
	return httpClient
}

func New(ctx context.Context, c *clientcredentials.Config) *http.Client {
	return oauth2.NewClient(ctx, c.TokenSource(ctx))
}

func OAuth2(u *url.URL, clientID, clientSecret string) *clientcredentials.Config {
	return &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     u.String() + "/oauth/token",
		EndpointParams: url.Values{
			"audience": {u.String() + "/api/v2/"},
		},
	}
}
