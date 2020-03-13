package management

import (
	"encoding/json"
	"reflect"
)

type Connection struct {
	// A generated string identifying the connection.
	ID *string `json:"id,omitempty"`

	// The name of the connection. Must start and end with an alphanumeric
	// character and can only contain alphanumeric characters and '-'. Max
	// length 128.
	Name *string `json:"name,omitempty"`

	// The identity provider identifier for the connection. Can be any of the
	// following:
	//
	// "ad", "adfs", "amazon", "dropbox", "bitbucket", "aol", "auth0-adldap",
	// "auth0-oidc", "auth0", "baidu", "bitly", "box", "custom", "daccount",
	// "dwolla", "email", "evernote-sandbox", "evernote", "exact", "facebook",
	// "fitbit", "flickr", "github", "google-apps", "google-oauth2", "guardian",
	//  "instagram", "ip", "linkedin", "miicard", "oauth1", "oauth2",
	// "office365", "paypal", "paypal-sandbox", "pingfederate",
	// "planningcenter", "renren", "salesforce-community", "salesforce-sandbox",
	//  "salesforce", "samlp", "sharepoint", "shopify", "sms", "soundcloud",
	// "thecity-sandbox", "thecity", "thirtysevensignals", "twitter", "untappd",
	//  "vkontakte", "waad", "weibo", "windowslive", "wordpress", "yahoo",
	// "yammer" or "yandex".
	Strategy *string `json:"strategy,omitempty"`

	// True if the connection is domain level
	IsDomainConnection *bool `json:"is_domain_connection,omitempty"`

	// Options for validation.
	Options    interface{}     `json:"-"`
	RawOptions json.RawMessage `json:"options,omitempty"`

	// The identifiers of the clients for which the connection is to be
	// enabled. If the array is empty or the property is not specified, no
	// clients are enabled.
	EnabledClients []interface{} `json:"enabled_clients,omitempty"`

	// Defines the realms for which the connection will be used (ie: email
	// domains). If the array is empty or the property is not specified, the
	// connection name will be added as realm.
	Realms []interface{} `json:"realms,omitempty"`

	Metadata *interface{} `json:"metadata,omitempty"`
}

func (c *Connection) MarshalJSON() ([]byte, error) {

	type connection Connection

	if c.Options != nil {
		b, err := json.Marshal(c.Options)
		if err != nil {
			return nil, err
		}
		c.RawOptions = b
	}

	return json.Marshal((*connection)(c))
}

func (c *Connection) UnmarshalJSON(b []byte) error {

	type connection Connection

	err := json.Unmarshal(b, (*connection)(c))
	if err != nil {
		return err
	}

	if c.Strategy != nil {

		var v interface{}

		switch *c.Strategy {
		case "auth0":
			v = &ConnectionOptions{}
		case "google-oauth2":
			v = &ConnectionOptionsGooleOAuth2{}
		case "facebook":
			v = &ConnectionOptionsFacebook{}
		case "apple":
			v = &ConnectionOptionsApple{}
		case "linkedin":
			v = &ConnectionOptionsLinkedin{}
		case "github":
			v = &ConnectionOptionsGitHub{}
		case "windowslive":
			v = &ConnectionOptionsWindowsLive{}
		case "salesforce":
			v = &ConnectionOptionsSalesforce{}
		case "email":
			v = &ConnectionOptionsEmail{}
		case "sms":
			v = &ConnectionOptionsSMS{}
		case "oidc":
			v = &ConnectionOptionsOIDC{}
		case "waad":
			v = &ConnectionOptionsAzureAD{}
		default:
			v = &map[string]interface{}{}
		}

		err = json.Unmarshal(c.RawOptions, &v)
		if err != nil {
			return err
		}

		c.Options = reflect.ValueOf(v).Elem().Interface()
	}

	return nil
}

type ConnectionOptions struct {

	// Options for multifactor authentication. Can be used to set active and
	// return_enroll_settings.
	MFA map[string]interface{} `json:"mfa,omitempty"`

	// Options for validation.
	Validation map[string]interface{} `json:"validation,omitempty"`

	// Password strength level, can be one of:
	// "none", "low", "fair", "good", "excellent" or null.
	PasswordPolicy *string `json:"passwordPolicy,omitempty"`

	// Options for password history policy.
	PasswordHistory map[string]interface{} `json:"password_history,omitempty"`

	// Options for password expiration policy.
	PasswordNoPersonalInfo map[string]interface{} `json:"password_no_personal_info,omitempty"`

	// Options for password dictionary policy.
	PasswordDictionary map[string]interface{} `json:"password_dictionary,omitempty"`

	// Options for password complexity options.
	PasswordComplexityOptions map[string]interface{} `json:"password_complexity_options,omitempty"`

	EnabledDatabaseCustomization *bool `json:"enabledDatabaseCustomization,omitempty"`

	BruteForceProtection *bool `json:"brute_force_protection,omitempty"`

	ImportMode *bool `json:"import_mode,omitempty"`

	DisableSignup *bool `json:"disable_signup,omitempty"`

	RequiresUsername *bool `json:"requires_username,omitempty"`

	// Scripts for the connection
	// Allowed keys are: "get_user", "login", "create", "verify", "change_password", "delete" or "change_email".
	CustomScripts map[string]interface{} `json:"customScripts,omitempty"`
	// configuration variables that can be used in custom scripts
	Configuration map[string]interface{} `json:"configuration,omitempty"`

	StrategyVersion *int `json:"strategy_version"`
}

type ConnectionOptionsGooleOAuth2 struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	AllowedAudiences []interface{} `json:"allowed_audiences,omitempty"`

	Email                  *bool `json:"email,omitempty"`
	Profile                *bool `json:"profile,omitempty"`
	Contacts               *bool `json:"contacts,omitempty"`
	Blogger                *bool `json:"blogger,omitempty"`
	Calendar               *bool `json:"calendar,omitempty"`
	Gmail                  *bool `json:"gmail,omitempty"`
	GooglePlus             *bool `json:"google_plus,omitempty"`
	Orkut                  *bool `json:"orkut,omitempty"`
	PicasaWeb              *bool `json:"picasa_web,omitempty"`
	Tasks                  *bool `json:"tasks,omitempty"`
	Youtube                *bool `json:"youtube,omitempty"`
	AdsenseManagement      *bool `json:"adsense_management,omitempty"`
	GoogleAffiliateNetwork *bool `json:"google_affiliate_network,omitempty"`
	Analytics              *bool `json:"analytics,omitempty"`
	GoogleBooks            *bool `json:"google_books,omitempty"`
	GoogleCloudStorage     *bool `json:"google_cloud_storage,omitempty"`
	ContentAPIForShopping  *bool `json:"content_api_for_shopping,omitempty"`
	ChromeWebStore         *bool `json:"chrome_web_store,omitempty"`
	DocumentList           *bool `json:"document_list,omitempty"`
	GoogleDrive            *bool `json:"google_drive,omitempty"`
	GoogleDriveFiles       *bool `json:"google_drive_files,omitempty"`
	LatitudeBest           *bool `json:"latitude_best,omitempty"`
	LatitudeCity           *bool `json:"latitude_city,omitempty"`
	Moderator              *bool `json:"moderator,omitempty"`
	Sites                  *bool `json:"sites,omitempty"`
	Spreadsheets           *bool `json:"spreadsheets,omitempty"`
	URLShortener           *bool `json:"url_shortener,omitempty"`
	WebmasterTools         *bool `json:"webmaster_tools,omitempty"`
	Coordinate             *bool `json:"coordinate,omitempty"`
	CoordinateReadonly     *bool `json:"coordinate_readonly,omitempty"`

	Scope []interface{} `json:"scope,omitempty"`
}

type ConnectionOptionsFacebook struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	Email                       *bool `json:"email,omitempty"`
	GroupsAccessMemberInfo      *bool `json:"groups_access_member_info,omitempty"`
	PublishToGroups             *bool `json:"publish_to_groups,omitempty"`
	UserAgeRange                *bool `json:"user_age_range,omitempty"`
	UserBirthday                *bool `json:"user_birthday,omitempty"`
	AdsManagement               *bool `json:"ads_management,omitempty"`
	AdsRead                     *bool `json:"ads_read,omitempty"`
	ReadAudienceNetworkInsights *bool `json:"read_audience_network_insights,omitempty"`
	ReadInsights                *bool `json:"read_insights,omitempty"`
	ManageNotifications         *bool `json:"manage_notifications,omitempty"`
	PublishActions              *bool `json:"publish_actions,omitempty"`
	ReadMailbox                 *bool `json:"read_mailbox,omitempty"`
	PublicProfile               *bool `json:"public_profile,omitempty"`
	UserEvents                  *bool `json:"user_events,omitempty"`
	UserFriends                 *bool `json:"user_friends,omitempty"`
	UserGender                  *bool `json:"user_gender,omitempty"`
	UserHometown                *bool `json:"user_hometown,omitempty"`
	UserLikes                   *bool `json:"user_likes,omitempty"`
	UserLink                    *bool `json:"user_link,omitempty"`
	UserLocation                *bool `json:"user_location,omitempty"`
	UserPhotos                  *bool `json:"user_photos,omitempty"`
	UserPosts                   *bool `json:"user_posts,omitempty"`
	UserTaggedPlaces            *bool `json:"user_tagged_places,omitempty"`
	UserVideos                  *bool `json:"user_videos,omitempty"`
	BusinessManagement          *bool `json:"business_management,omitempty"`
	LeadsRetrieval              *bool `json:"leads_retrieval,omitempty"`
	ManagePages                 *bool `json:"manage_pages,omitempty"`
	PagesManageCTA              *bool `json:"pages_manage_cta,omitempty"`
	PagesManageInstantArticles  *bool `json:"pages_manage_instant_articles,omitempty"`
	PagesShowList               *bool `json:"pages_show_list,omitempty"`
	PagesMessaging              *bool `json:"pages_messaging,omitempty"`
	PagesMessagingPhoneNumber   *bool `json:"pages_messaging_phone_number,omitempty"`
	PagesMessagingSubscriptions *bool `json:"pages_messaging_subscriptions,omitempty"`
	PublishPages                *bool `json:"publish_pages,omitempty"`
	PublishVideo                *bool `json:"publish_video,omitempty"`
	ReadPageMailboxes           *bool `json:"read_page_mailboxes,omitempty"`
	ReadStream                  *bool `json:"read_stream,omitempty"`
	UserGroups                  *bool `json:"user_groups,omitempty"`
	UserManagedGroups           *bool `json:"user_managed_groups,omitempty"`
	UserStatus                  *bool `json:"user_status,omitempty"`
	AllowContextProfileField    *bool `json:"allow_context_profile_field,omitempty"`

	// Scope is a comma separated list of scopes.
	Scope *string `json:"scope,omitempty"`
}

type ConnectionOptionsApple struct {
	ClientID               *string `json:"client_id,omitempty"`
	ClientSecretSigningKey *string `json:"app_secret,omitempty"`

	TeamID *string `json:"team_id,omitempty"`
	KeyID  *string `json:"kid,omitempty"`

	Name  *bool `json:"name,omitempty"`
	Email *bool `json:"email,omitempty"`

	Scope *string `json:"scope,omitempty"`
}

type ConnectionOptionsLinkedin struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	StrategyVersion *int `json:"strategy_version"`

	Email        *bool `json:"email,omitempty"`
	Profile      *bool `json:"profile,omitempty"`
	BasicProfile *bool `json:"basic_profile,omitempty"`

	Scope []interface{} `json:"scope,omitempty"`

	SetUserAttributes *string `json:"set_user_root_attributes,omitempty"`
}

type ConnectionOptionsGitHub struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	Email          *bool `json:"email,omitempty"`
	ReadUser       *bool `json:"read_user,omitempty"`
	Follow         *bool `json:"follow,omitempty"`
	PublicRepo     *bool `json:"public_repo,omitempty"`
	Repo           *bool `json:"repo,omitempty"`
	RepoDeployment *bool `json:"repo_deployment,omitempty"`
	RepoStatus     *bool `json:"repo_status,omitempty"`
	DeleteRepo     *bool `json:"delete_repo,omitempty"`
	Notifications  *bool `json:"notifications,omitempty"`
	Gist           *bool `json:"gist,omitempty"`
	ReadRepoHook   *bool `json:"read_repo_hook,omitempty"`
	WriteRepoHook  *bool `json:"write_repo_hook,omitempty"`
	AdminRepoHook  *bool `json:"admin_repo_hook,omitempty"`
	ReadOrg        *bool `json:"read_org,omitempty"`
	AdminOrg       *bool `json:"admin_org,omitempty"`
	ReadPublicKey  *bool `json:"read_public_key,omitempty"`
	WritePublicKey *bool `json:"write_public_key,omitempty"`
	AdminPublicKey *bool `json:"admin_public_key,omitempty"`
	WriteOrg       *bool `json:"write_org,omitempty"`
	Profile        *bool `json:"profile,omitempty"`

	Scope []interface{} `json:"scope,omitempty"`

	SetUserAttributes *string `json:"set_user_root_attributes,omitempty"`
}

type ConnectionOptionsEmail struct {
	Name  *string                         `json:"name,omitempty"`
	Email *ConnectionOptionsEmailSettings `json:"email,omitempty"`
	OTP   *ConnectionOptionsOTP           `json:"totp,omitempty"`

	DisableSignup        *bool `json:"disable_signup,omitempty"`
	BruteForceProtection *bool `json:"brute_force_protection,omitempty"`
}

type ConnectionOptionsEmailSettings struct {
	Syntax  *string `json:"syntax,omitempty"`
	From    *string `json:"from,omitempty"`
	Subject *string `json:"subject,omitempty"`
	Body    *string `json:"body,omitempty"`
}

type ConnectionOptionsOTP struct {
	TimeStep *int `json:"time_step,omitempty"`
	Length   *int `json:"length,omitempty"`
}

type ConnectionOptionsSMS struct {
	Name     *string `json:"name,omitempty"`
	From     *string `json:"from,omitempty"`
	Syntax   *string `json:"syntax,omitempty"`
	Template *string `json:"template,omitempty"`

	OTP *ConnectionOptionsOTP `json:"totp,omitempty"`

	TwilioSID           *string `json:"twilio_sid"`
	TwilioToken         *string `json:"twilio_token"`
	MessagingServiceSID *string `json:"messaging_service_sid"`

	DisableSignup        *bool `json:"disable_signup,omitempty"`
	BruteForceProtection *bool `json:"brute_force_protection,omitempty"`
}

type ConnectionOptionsWindowsLive struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	StrategyVersion *int `json:"strategy_version"`

	OfflineAccess   *bool `json:"offline_access,omitempty"`
	UserUpdate      *bool `json:"graph_user_update,omitempty"`
	UserActivity    *bool `json:"graph_user_activity,omitempty"`
	Device          *bool `json:"graph_device,omitempty"`
	Emails          *bool `json:"graph_emails,omitempty"`
	NotesUpdate     *bool `json:"graph_notes_update,omitempty"`
	User            *bool `json:"graph_user,omitempty"`
	DeviceCommand   *bool `json:"graph_device_command,omitempty"`
	EmailsUpdate    *bool `json:"graph_emails_update,omitempty"`
	Calendars       *bool `json:"graph_calendars,omitempty"`
	CalendarsUpdate *bool `json:"graph_calendars_update,omitempty"`
	Contacts        *bool `json:"graph_contacts,omitempty"`
	ContactsUpdate  *bool `json:"graph_contacts_update,omitempty"`
	Files           *bool `json:"graph_files,omitempty"`
	FilesAll        *bool `json:"graph_files_all,omitempty"`
	FilesUpdate     *bool `json:"graph_files_update,omitempty"`
	FilesAllUpdate  *bool `json:"graph_files_all_update,omitempty"`
	Notes           *bool `json:"graph_notes,omitempty"`
	NotesCreate     *bool `json:"graph_notes_create,omitempty"`
	Tasks           *bool `json:"graph_tasks,omitempty"`
	TasksUpdate     *bool `json:"graph_tasks_update,omitempty"`
	Signin          *bool `json:"signin,omitempty"`

	Scope []interface{} `json:"scope,omitempty"`

	SetUserAttributes *string `json:"set_user_root_attributes,omitempty"`
}

// Salesforce
type ConnectionOptionsSalesforce struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	Profile *bool `json:"profile,omitempty"`

	Scope []interface{} `json:"scope,omitempty"`

	CommunityBaseURL  *string `json:"community_base_url,omitempty"`
	SetUserAttributes *string `json:"set_user_root_attributes,omitempty"`
}

type ConnectionOptionsOIDC struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	TenantDomain  *string       `json:"tenant_domain,omitempty"`
	DomainAliases []interface{} `json:"domain_aliases,omitempty"`
	LogoURL       *string       `json:"icon_url,omitempty"`

	DiscoveryURL          *string `json:"discovery_url"`
	AuthorizationEndpoint *string `json:"authorization_endpoint"`
	Issuer                *string `json:"issuer"`
	JWKSURI               *string `json:"jwks_uri"`
	Type                  *string `json:"type"`
	UserInfoEndpoint      *string `json:"userinfo_endpoint"`
	TokenEndpoint         *string `json:"token_endpoint"`

	Scope *string `json:"scope,omitempty"`

	// "scope": "openid profile email",
	// "icon_url": "https://alexkappa.com/logo.png",
	// "discovery_url": "https://alexkappa.eu.auth0.com",
	// "authorization_endpoint": "https://alexkappa.eu.auth0.com/authorize",
	// "issuer": "https://alexkappa.eu.auth0.com/",
	// "jwks_uri": "https://alexkappa.eu.auth0.com/.well-known/jwks.json",
	// "type": "front_channel",
	// "userinfo_endpoint": "https://alexkappa.eu.auth0.com/userinfo",
	// "token_endpoint": null,
	// "client_id": "foo",
	// "domain_aliases": [
	// 	"alexkappa.com",
	// 	"api.alexkappa.com"
	// ],
	// "tenant_domain": "alexkappa.com"
}

// Azure AD
type ConnectionOptionsAzureAD struct {
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`

	TenantDomain  *string       `json:"tenant_domain,omitempty"`
	Domain        *string       `json:"domain,omitempty"`
	DomainAliases []interface{} `json:"domain_aliases,omitempty"`
	LogoURL       *string       `json:"icon_url,omitempty"`

	IdentityAPI *string `json:"identity_api"`

	WAADProtocol       *string `json:"waad_protocol,omitempty"`
	WAADCommonEndpoint *bool   `json:"waad_common_endpoint,omitempty"`

	UseWSFederation     *bool   `json:"use_wsfed,omitempty"`
	UseCommonEndpoint   *bool   `json:"useCommonEndpoint,omitempty"`
	EnableUsersAPI      *bool   `json:"api_enable_users,omitempty"`
	MaxGroupsToRetrieve *string `json:"max_groups_to_retrieve,omitempty"`

	BasicProfile     *bool `json:"basic_profile,omitempty"`
	ExtProfile       *bool `json:"ext_profile,omitempty"`
	ExtGroups        *bool `json:"ext_groups,omitempty"`
	ExtNestedGroups  *bool `json:"ext_nested_groups,omitempty"`
	ExtAdmin         *bool `json:"ext_admin,omitempty"`
	ExtIsSuspended   *bool `json:"ext_is_suspended,omitempty"`
	ExtAgreedTerms   *bool `json:"ext_agreed_terms,omitempty"`
	ExtAssignedPlans *bool `json:"ext_assigned_plans,omitempty"`
}

type ConnectionManager struct {
	*Management
}

type ConnectionOptionsTotp struct {
	TimeStep *int `json:"time_step,omitempty"`
	Length   *int `json:"length,omitempty"`
}

type ConnectionList struct {
	List
	Connections []*Connection `json:"connections"`
}

func newConnectionManager(m *Management) *ConnectionManager {
	return &ConnectionManager{m}
}

// Create a new connection.
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/post_connections
func (m *ConnectionManager) Create(c *Connection) error {
	return m.post(m.uri("connections"), c)
}

// Read retrieves a connection by its id.
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/get_connections_by_id
func (m *ConnectionManager) Read(id string) (c *Connection, err error) {
	err = m.get(m.uri("connections", id), &c)
	return
}

// List all connections.
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/get_connections
func (m *ConnectionManager) List(opts ...ListOption) (c *ConnectionList, err error) {
	opts = m.defaults(opts)
	err = m.get(m.uri("connections")+m.q(opts), &c)
	return
}

// Update a connection.
//
// Note: if you use the options parameter, the whole options object will be
// overridden, so ensure that all parameters are present.
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/patch_connections_by_id
func (m *ConnectionManager) Update(id string, c *Connection) (err error) {
	return m.patch(m.uri("connections", id), c)
}

// Delete a connection and all its users.
//
// See: https://auth0.com/docs/api/management/v2#!/Connections/delete_connections_by_id
func (m *ConnectionManager) Delete(id string) (err error) {
	return m.delete(m.uri("connections", id))
}

// ReadByName retrieves a connection by its name. This is a helper method when a
// connection id is not readily available.
func (m *ConnectionManager) ReadByName(name string) (*Connection, error) {
	c, err := m.List(Parameter("name", name))
	if err != nil {
		return nil, err
	}
	if len(c.Connections) > 0 {
		return c.Connections[0], nil
	}
	return nil, &managementError{404, "Not Found", "Connection not found"}
}
