package management

import "time"

type Identity struct {
	Connection *string `json:"connection,omitempty"`
	UserID     *string `json:"user_id,omitempty"`
	Provider   *string `json:"provider,omitempty"`
	IsSocial   *bool   `json:"isSocial,omitempty"`
}

type User struct {

	// The users identifier.
	ID *string `json:"user_id,omitempty"`

	// The connection the user belongs to.
	Connection *string `json:"connection,omitempty"`

	// The user's email
	Email *string `json:"email,omitempty"`

	// The users name
	Name *string `json:"name,omitempty"`

	// The users given name
	GivenName *string `json:"given_name,omitempty"`

	// The users family name
	FamilyName *string `json:"family_name,omitempty"`

	// The user's username. Only valid if the connection requires a username
	Username *string `json:"username,omitempty"`

	// The user's nickname
	Nickname *string `json:"nickname,omitempty"`

	// The user's password (mandatory for non SMS connections)
	Password *string `json:"password,omitempty"`

	// The user's phone number (following the E.164 recommendation), only valid
	// for users to be added to SMS connections.
	PhoneNumber *string `json:"phone_number,omitempty"`

	// The time the user is created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The last time the user is updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// The last time the user has logged in.
	LastLogin *time.Time `json:"last_login,omitempty"`

	// UserMetadata holds data that the user has read/write access to (e.g.
	// color_preference, blog_url, etc).
	UserMetadata map[string]interface{} `json:"user_metadata,omitempty"`

	Identities []*Identity `json:"identities,omitempty"`

	// True if the user's email is verified, false otherwise. If it is true then
	// the user will not receive a verification email, unless verify_email: true
	// was specified.
	EmailVerified *bool `json:"email_verified,omitempty"`

	// If true, the user will receive a verification email after creation, even
	// if created with email_verified set to true. If false, the user will not
	// receive a verification email, even if created with email_verified set to
	// false. If unspecified, defaults to the behavior determined by the value
	// of email_verified.
	VerifyEmail *bool `json:"verify_email,omitempty"`

	// True if the user's phone number is verified, false otherwise. When the
	// user is added to a SMS connection, they will not receive an verification
	// SMS if this is true.
	PhoneVerified *bool `json:"phone_verified,omitempty"`

	// AppMetadata holds data that the user has read-only access to (e.g. roles,
	// permissions, vip, etc).
	AppMetadata map[string]interface{} `json:"app_metadata,omitempty"`

	// The user's picture url
	Picture *string `json:"picture,omitempty"`

	// True if the user is blocked from the application, false if the user is enabled
	Blocked *bool `json:"blocked,omitempty"`
}

func (u *User) String() string {
	return Stringify(u)
}

type UserManager struct {
	m *Management
}

func NewUserManager(m *Management) *UserManager {
	return &UserManager{m}
}

// Creates a new user. It works only for database and passwordless connections.
//
// The samples on the right show you every attribute that could be used. The
// attribute connection is always mandatory but depending on the type of
// connection you are using there could be others too. For instance, database
// connections require `email` and `password`.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/post_users
func (um *UserManager) Create(u *User) error {
	return um.m.post(um.m.uri("users"), u)
}

// This endpoint can be used to retrieve user details given the user_id.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/get_users_by_id
func (um *UserManager) Read(id string, opts ...ReqOption) (*User, error) {
	u := new(User)
	err := um.m.get(um.m.uri("users", id)+um.m.q(opts), u)
	return u, err
}

// Updates a user.
//
// The following attributes can be updated at the root level:
//
// - `app_metadata`
// - `blocked`
// - `email`
// - `email_verified`
// - `family_name`
// - `given_name`
// - `name`
// - `nickname`
// - `password`
// - `phone_number`
// - `phone_verified`
// - `picture`
// - `username`
// - `user_metadata`
// - `verify_email`
//
// See: https://auth0.com/docs/api/management/v2#!/Users/patch_users_by_id
func (um *UserManager) Update(id string, u *User) (err error) {
	return um.m.patch(um.m.uri("users", id), u)
}

// This endpoint can be used to delete a single user based on the id.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/delete_users_by_id
func (um *UserManager) Delete(id string) (err error) {
	return um.m.delete(um.m.uri("users", id))
}

// This endpoint can be used to retrieve a list of users.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/get_users
func (um *UserManager) List(opts ...ReqOption) (us []*User, err error) {
	err = um.m.get(um.m.uri("users")+um.m.q(opts), &us)
	return
}

// Search is an alias for List.
func (um *UserManager) Search(opts ...ReqOption) (us []*User, err error) {
	return um.List(opts...)
}

// If Auth0 is the identify provider (idP), the email address associated with a
// user is saved in lower case, regardless of how you initially provided it.
// For example, if you register a user as JohnSmith@example.com, Auth0 saves the
// user's email as johnsmith@example.com.
//
// In cases where Auth0 is not the idP, the `email` is stored based on the rules
// of idP, so make sure the search is made using the correct capitalization.
//
// When using this endpoint, make sure that you are searching for users via
// email addresses using the correct case.
//
// See: https://auth0.com/docs/api/management/v2#!/Users_By_Email/get_users_by_email
func (um *UserManager) ListByEmail(email string, opts ...ReqOption) (us []*User, err error) {
	opts = append(opts, Parameter("email", email))
	err = um.m.get(um.m.uri("users-by-email")+um.m.q(opts), &us)
	return
}

// List the the roles associated with a user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/get_user_roles
func (um *UserManager) GetRoles(id string, opts ...ReqOption) (roles []*Role, err error) {
	err = um.m.get(um.m.uri("users", id, "roles")+um.m.q(opts), &roles)
	return roles, err
}

// Assign roles to a user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/post_user_roles
func (um *UserManager) AssignRoles(id string, roles ...*Role) error {
	r := make(map[string][]*string)
	r["roles"] = make([]*string, len(roles))
	for i, role := range roles {
		r["roles"][i] = role.ID
	}
	return um.m.post(um.m.uri("users", id, "roles"), &r)
}

// Removes roles from a user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/delete_user_roles
func (um *UserManager) RemoveRoles(id string, roles ...*Role) error {
	r := make(map[string][]*string)
	r["roles"] = make([]*string, len(roles))
	for i, role := range roles {
		r["roles"][i] = role.ID
	}
	return um.m.request("DELETE", um.m.uri("users", id, "roles"), &r)
}

// List the permissions associated to the user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/get_permissions
func (um *UserManager) Permissions(id string, opts ...ReqOption) (permissions []*Permission, err error) {
	err = um.m.get(um.m.uri("users", id, "permissions")+um.m.q(opts), &permissions)
	return permissions, err
}

// Assign permissions to the user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/post_permissions
func (um *UserManager) AssignPermissions(id string, permissions ...*Permission) error {
	p := make(map[string][]*Permission)
	p["permissions"] = permissions
	return um.m.post(um.m.uri("users", id, "permissions"), &p)
}

// Removes permissions from a user.
//
// See: https://auth0.com/docs/api/management/v2#!/Users/delete_permissions
func (um *UserManager) RemovePermissions(id string, permissions ...*Permission) error {
	p := make(map[string][]*Permission)
	p["permissions"] = permissions
	return um.m.request("DELETE", um.m.uri("users", id, "permissions"), &p)
}
