package management

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"testing"
	"time"
)

var m *Management

var (
	domain       = os.Getenv("AUTH0_DOMAIN")
	clientID     = os.Getenv("AUTH0_CLIENT_ID")
	clientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
)

func init() {
	var err error
	m, err = New(domain, clientID, clientSecret, WithDebug(false))
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

func TestWithHTTPClient(t *testing.T) {
	var err error

	timeout := time.Hour * 42
	expected := &timeout
	client := &http.Client{Timeout: timeout}

	m, err = New(domain, clientID, clientSecret, WithHTTPClient(client))
	if err != nil {
		panic(err)
	}

	if m.http.Timeout != *expected {
		t.Errorf("Expected %q, but got %q", expected, m.http.Timeout)
	}

	client = &http.Client{}

	m, err = New(domain, clientID, clientSecret, WithHTTPClient(client))
	if err != nil {
		panic(err)
	}

	if m.http.Timeout != 0 {
		t.Errorf("Expected %q, but got %q", 0, m.http.Timeout)
	}

	jar, _ := cookiejar.New(&cookiejar.Options{})
	client = &http.Client{Jar: jar}
	m, err = New(domain, clientID, clientSecret, WithHTTPClient(client))
	if err != nil {
		panic(err)
	}

	if m.http.Jar != jar {
		t.Errorf("Expected %#v, but got %#v", jar, m.http.Jar)
	}
}
