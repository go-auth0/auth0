package management

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"
)

var m *Management

var (
	domain       = os.Getenv("AUTH0_DOMAIN")
	clientID     = os.Getenv("AUTH0_CLIENT_ID")
	clientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	debug        = os.Getenv("AUTH0_DEBUG")
)

func init() {
	var err error
	m, err = New(domain, clientID, clientSecret,
		WithDebug(debug == "true" || debug == "1" || debug == "on"))
	if err != nil {
		panic(err)
	}
}

func TestNew(t *testing.T) {
	for _, domain := range []string{
		"example.com ",
		" example.com",
		" example.com ",
		"%2Fexample.com",
		" a.b.c.example.com",
	} {
		_, err := New(domain, "", "")
		if err == nil {
			t.Errorf("expected New to fail with domain %q", domain)
		}
	}
}

func TestOptionFields(t *testing.T) {
	v := make(url.Values)
	WithFields("foo", "bar")(v)

	fields := v.Get("fields")
	if fields != "foo,bar" {
		t.Errorf("Expected %q, but got %q", fields, "foo,bar")
	}

	includeFields := v.Get("include_fields")
	if includeFields != "true" {
		t.Errorf("Expected %q, but got %q", includeFields, "true")
	}

	WithoutFields("foo", "bar")(v)

	includeFields = v.Get("include_fields")
	if includeFields != "false" {
		t.Errorf("Expected %q, but got %q", includeFields, "true")
	}
}

func TestOptionPage(t *testing.T) {
	v := make(url.Values)
	Page(3)(v)
	PerPage(10)(v)

	page := v.Get("page")
	if page != "3" {
		t.Errorf("Expected %q, but got %q", page, "3")
	}

	perPage := v.Get("per_page")
	if perPage != "10" {
		t.Errorf("Expected %q, but got %q", perPage, "3")
	}
}

func TestOptionTotals(t *testing.T) {
	v := make(url.Values)
	IncludeTotals(true)(v)

	includeTotals := v.Get("include_totals")
	if includeTotals != "true" {
		t.Errorf("Expected %q, but got %q", includeTotals, "true")
	}
}

func TestOptionParameter(t *testing.T) {
	v := make(url.Values)
	Parameter("foo", "123")(v)
	Parameter("bar", "xyz")(v)

	foo := v.Get("foo")
	if foo != "123" {
		t.Errorf("Expected %q, but got %q", foo, "123")
	}

	bar := v.Get("bar")
	if bar != "xyz" {
		t.Errorf("Expected %q, but got %q", bar, "xyz")
	}
}

func TestStringify(t *testing.T) {

	expected := `{
  "foo": "bar"
}`

	v := struct {
		Foo string `json:"foo"`
	}{
		"bar",
	}

	s := Stringify(v)

	if s != expected {
		t.Errorf("Expected %q, but got %q", expected, s)
	}
}

func TestRequest_RateLimited(t *testing.T) {
	expectedReset := time.Now()
	expectedLimit := 2
	expectedRemaining := 3

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set(RateLimitReset, strconv.Itoa(int(expectedReset.Unix())))
		rw.Header().Set(RateLimitLimit, strconv.Itoa(expectedLimit))
		rw.Header().Set(RateLimitRemaining, strconv.Itoa(expectedRemaining))
		rw.WriteHeader(http.StatusTooManyRequests)
		rw.Write([]byte(`{"statusCode": 429, "error": "too many requests", "message": "get outta here"}`))
	}))
	defer server.Close()

	m := &Management{
		timeout: time.Second,
		ctx:     context.Background(),
		http:    http.DefaultClient,
	}

	err := m.request("GET", server.URL, nil)
	var rateLimitErr *RateLimitError
	if !errors.As(err, &rateLimitErr) {
		t.Errorf("expected RateLimitError, but got %s", err)
	}

	actualLimit := rateLimitErr.RateLimit.Limit
	if int(actualLimit) != expectedLimit {
		t.Errorf("expected Limit to be %d, got %d", expectedLimit, actualLimit)
	}
	actualReset := rateLimitErr.RateLimit.Reset
	if actualReset.Unix() != expectedReset.Unix() {
		t.Errorf("expected Reset to be %d, got %d", expectedReset.Unix(), actualReset.Unix())
	}
	actualRemaining := rateLimitErr.RateLimit.Remaining
	if int(actualRemaining) != expectedRemaining {
		t.Errorf("expected Remaining to be %d, got %d", expectedRemaining, actualRemaining)
	}
}
