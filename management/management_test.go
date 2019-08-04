package management

import (
	"net/url"
	"os"
	"testing"
)

var m *Management

var (
	auth0Domain       = os.Getenv("AUTH0_DOMAIN")
	auth0ClientID     = os.Getenv("AUTH0_CLIENT_ID")
	auth0ClientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
)

func init() {
	var err error
	m, err = New(auth0Domain, auth0ClientID, auth0ClientSecret, WithDebug(true))
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
