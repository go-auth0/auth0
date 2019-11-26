package management

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"gopkg.in/auth0.v1/internal/client"
)

// Management is an Auth0 management client used to interact with the Auth0
// Management API v2.
//
type Management struct {
	// Client manages Auth0 Client (also known as Application) resources.
	Client *ClientManager

	// ClientGrant manages Auth0 ClientGrant resources.
	ClientGrant *ClientGrantManager

	// ResourceServer manages Auth0 Resource Server (also known as API)
	// resources.
	ResourceServer *ResourceServerManager

	// Connection manages Auth0 Connection resources.
	Connection *ConnectionManager

	// CustomDomain manages Auth0 Custom Domains.
	CustomDomain *CustomDomainManager

	// Grant manages Auth0 Grants.
	Grant *GrantManager

	// Log reads Auth0 Logs.
	Log *LogManager

	// RoleManager manages Auth0 Roles.
	Role *RoleManager

	// RuleManager manages Auth0 Rules.
	Rule *RuleManager

	// RuleManager manages Auth0 Rule Configurations.
	RuleConfig *RuleConfigManager

	// Email manages Auth0 Email Providers.
	Email *EmailManager

	// EmailTemplate manages Auth0 Email Templates.
	EmailTemplate *EmailTemplateManager

	// User manages Auth0 User resources.
	User *UserManager

	// Job manages Auth0 jobs.
	Job *JobManager

	// Tenant manages your Auth0 Tenant.
	Tenant *TenantManager

	// Ticket creates verify email or change password tickets.
	Ticket *TicketManager

	// Stat is used to retrieve usage statistics.
	Stat *StatManager

	// Branding settings such as company logo or primary color.
	Branding *BrandingManager

	// Guardian manages your Auth0 Guardian settings
	Guardian *GuardianManager

	// Prompt manages your prompt settings.
	Prompt *PromptManager

	// Blacklist manages the auth0 blacklists
	Blacklist *BlacklistManager

	url      *url.URL
	basePath string
	timeout  time.Duration
	debug    bool
	ctx      context.Context

	http *http.Client
}

// New creates a new Auth0 Management client by authenticating using the
// supplied client id and secret.
func New(domain, clientID, clientSecret string, options ...apiOption) (*Management, error) {

	// Ignore the scheme if it was defined in the domain variable. Then prefix
	// with https as its the only scheme supported by the Auth0 API.
	if i := strings.Index(domain, "//"); i != -1 {
		domain = domain[i+2:]
	}
	domain = "https://" + domain

	u, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}

	m := &Management{
		url:      u,
		basePath: "api/v2",
		timeout:  1 * time.Minute,
		debug:    false,
		ctx:      context.Background(),
	}

	for _, option := range options {
		option(m)
	}

	oauth2 := client.OAuth2(m.url, clientID, clientSecret)

	_, err = oauth2.Token(m.ctx)
	if err != nil {
		return nil, err
	}

	m.http = client.New(m.ctx, oauth2)
	m.http = client.WrapUserAgent(m.http)
	m.http = client.WrapRateLimit(m.http)
	if m.debug {
		m.http = client.WrapDebug(m.http)
	}

	m.Client = NewClientManager(m)
	m.ClientGrant = NewClientGrantManager(m)
	m.Connection = NewConnectionManager(m)
	m.CustomDomain = NewCustomDomainManager(m)
	m.Grant = NewGrantManager(m)
	m.Log = NewLogManager(m)
	m.ResourceServer = NewResourceServerManager(m)
	m.Role = NewRoleManager(m)
	m.Rule = NewRuleManager(m)
	m.RuleConfig = NewRuleConfigManager(m)
	m.EmailTemplate = NewEmailTemplateManager(m)
	m.Email = NewEmailManager(m)
	m.User = NewUserManager(m)
	m.Job = NewJobManager(m)
	m.Tenant = NewTenantManager(m)
	m.Ticket = NewTicketManager(m)
	m.Stat = NewStatManager(m)
	m.Branding = NewBrandingManager(m)
	m.Guardian = NewGuardianManager(m)
	m.Prompt = NewPromptManager(m)
	m.Blacklist = NewBlacklistManager(m)

	return m, nil
}

func (m *Management) uri(path ...string) string {
	return (&url.URL{
		Scheme: m.url.Scheme,
		Host:   m.url.Host,
		Path:   m.basePath + "/" + strings.Join(path, "/"),
	}).String()
}

func (m *Management) q(options []reqOption) string {
	if len(options) == 0 {
		return ""
	}
	v := make(url.Values)
	for _, option := range options {
		option(v)
	}
	return "?" + v.Encode()
}

func (m *Management) request(method, uri string, v interface{}) error {

	var payload bytes.Buffer
	if v != nil {
		err := json.NewEncoder(&payload).Encode(v)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, uri, &payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(m.ctx, m.timeout)
	defer cancel()

	if m.http == nil {
		m.http = http.DefaultClient
	}

	res, err := m.http.Do(req.WithContext(ctx))
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return err
		}
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return newError(res.Body)
	}

	if res.StatusCode != http.StatusNoContent {
		err := json.NewDecoder(res.Body).Decode(v)
		if err != nil {
			return err
		}
		return res.Body.Close()
	}

	return nil
}

func (m *Management) get(uri string, v interface{}) error {
	return m.request("GET", uri, v)
}

func (m *Management) post(uri string, v interface{}) error {
	return m.request("POST", uri, v)
}

func (m *Management) put(uri string, v interface{}) error {
	return m.request("PUT", uri, v)
}

func (m *Management) patch(uri string, v interface{}) error {
	return m.request("PATCH", uri, v)
}

func (m *Management) delete(uri string) error {
	return m.request("DELETE", uri, nil)
}

type apiOption func(*Management)

// WithTimeout configures the management client with a request timeout.
func WithTimeout(t time.Duration) apiOption {
	return func(m *Management) {
		m.timeout = t
	}
}

// WithDebug configures the management client to dump http requests and
// responses to stdout.
func WithDebug(d bool) apiOption {
	return func(m *Management) {
		m.debug = d
	}
}

// WitContext configures the management client to use the provided context
// instead of the provided one.
func WithContext(ctx context.Context) apiOption {
	return func(m *Management) {
		m.ctx = ctx
	}
}

type Error interface {
	Status() int
	error
}

type managementError struct {
	StatusCode int    `json:"statusCode"`
	Err        string `json:"error"`
	Message    string `json:"message"`
}

func newError(r io.Reader) error {
	m := &managementError{}
	err := json.NewDecoder(r).Decode(m)
	if err != nil {
		return err
	}
	return m
}

func (m *managementError) Error() string {
	return fmt.Sprintf("%d %s: %s", m.StatusCode, m.Err, m.Message)
}

func (m *managementError) Status() int {
	return m.StatusCode
}

// reqOption configures a call (typically to retrieve a resource) to Auth0 with
// query parameters.
type reqOption func(url.Values)

// WithFields configures a call to include the desired fields.
func WithFields(fields ...string) reqOption {
	return func(v url.Values) {
		v.Set("fields", strings.Join(fields, ","))
		v.Set("include_fields", "true")
	}
}

// WithoutFields configures a call to exclude the desired fields.
func WithoutFields(fields ...string) reqOption {
	return func(v url.Values) {
		v.Set("fields", strings.Join(fields, ","))
		v.Set("include_fields", "false")
	}
}

// Page configures a call to receive a specific page, if the results where
// concatenated.
func Page(page int) reqOption {
	return func(v url.Values) {
		v.Set("page", strconv.FormatInt(int64(page), 10))
	}
}

// PerPage configures a call to limit the amount of items in the result.
func PerPage(items int) reqOption {
	return func(v url.Values) {
		v.Set("per_page", strconv.FormatInt(int64(items), 10))
	}
}

// IncludeTotals configures a call to include totals.
func IncludeTotals(include bool) reqOption {
	return func(v url.Values) {
		v.Set("include_totals", strconv.FormatBool(include))
	}
}

// Parameter is a generic configuration to add arbitrary query parameters to
// calls made to Auth0.
func Parameter(key, value string) reqOption {
	return func(v url.Values) {
		v.Set(key, value)
	}
}
