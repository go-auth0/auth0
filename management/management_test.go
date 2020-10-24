package management

import (
	"net/http"
	"os"
	"testing"
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
	r, _ := http.NewRequest("GET", "/", nil)
	WithFields("foo", "bar").apply(r)

	v := r.URL.Query()

	fields := v.Get("fields")
	if fields != "foo,bar" {
		t.Errorf("Expected %q, but got %q", fields, "foo,bar")
	}

	includeFields := v.Get("include_fields")
	if includeFields != "true" {
		t.Errorf("Expected %q, but got %q", includeFields, "true")
	}

	WithoutFields("foo", "bar").apply(r)

	includeFields = v.Get("include_fields")
	if includeFields != "true" {
		t.Errorf("Expected %q, but got %q", includeFields, "true")
	}
}

func TestOptionPage(t *testing.T) {

	r, _ := http.NewRequest("GET", "/", nil)

	Page(3).apply(r)
	PerPage(10).apply(r)

	v := r.URL.Query()

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

	r, _ := http.NewRequest("GET", "/", nil)

	IncludeTotals(true).apply(r)

	v := r.URL.Query()

	includeTotals := v.Get("include_totals")
	if includeTotals != "true" {
		t.Errorf("Expected %q, but got %q", includeTotals, "true")
	}
}

func TestOptionParameter(t *testing.T) {

	r, _ := http.NewRequest("GET", "/", nil)

	Parameter("foo", "123").apply(r)
	Parameter("bar", "xyz").apply(r)

	v := r.URL.Query()

	foo := v.Get("foo")
	if foo != "123" {
		t.Errorf("Expected %q, but got %q", foo, "123")
	}

	bar := v.Get("bar")
	if bar != "xyz" {
		t.Errorf("Expected %q, but got %q", bar, "xyz")
	}
}

func TestOptionDefauls(t *testing.T) {

	r, _ := http.NewRequest("GET", "/", nil)

	applyListDefaults([]Option{
		PerPage(20),          // should be persist (default is 50)
		IncludeTotals(false), // should be altered to true by withListDefaults
	}).apply(r)

	v := r.URL.Query()

	perPage := v.Get("per_page")
	if perPage != "20" {
		t.Errorf("Expected %q, but got %q", perPage, "20")
	}

	includeTotals := v.Get("include_totals")
	if includeTotals != "true" {
		t.Errorf("Expected %q, but got %q", includeTotals, "true")
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
