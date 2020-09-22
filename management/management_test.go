package management

import (
	"context"
	"errors"
	"net/url"
	"os"
	"reflect"
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

func TestWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	n := m.WithContext(ctx)
	if n == m {
		t.Fatal("WithContext must return new instance")
	}

	v := reflect.ValueOf(n).Elem()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { // skip unexported fields
			fieldName := v.Field(i).Elem().Type().Name()
			if fieldName != "GuardianManager" { // skip more complicated manager
				n0 := v.Field(i).Elem().Field(0)
				if n != n0.Interface() {
					t.Fatalf("Field %s expected to point to new management struct", fieldName)
				}
			}
		}
	}

	_, err := n.User.List()
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected err to be context.Canceled, got %v", err)
	}
}
