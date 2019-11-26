package management

type Role struct {
	// A unique ID for the role.
	ID *string `json:"id,omitempty"`

	// The name of the role created.
	Name *string `json:"name,omitempty"`

	// A description of the role created.
	Description *string `json:"description,omitempty"`
}

type Permission struct {
	// The resource server that the permission is attached to.
	ResourceServerIdentifier *string `json:"resource_server_identifier,omitempty"`

	// The name of the resource server.
	ResourceServerName *string `json:"resource_server_name,omitempty"`

	// The name of the permission.
	Name *string `json:"permission_name,omitempty"`

	// The description of the permission.
	Description *string `json:"description,omitempty"`
}

func (r *Role) String() string {
	return Stringify(r)
}

type RoleManager struct {
	m *Management
}

func NewRoleManager(m *Management) *RoleManager {
	return &RoleManager{m}
}

// Create a new role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/post_roles
func (rm *RoleManager) Create(r *Role) error {
	return rm.m.post(rm.m.uri("roles"), r)
}

// Retrieve a role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/get_roles_by_id
func (rm *RoleManager) Read(id string, opts ...reqOption) (*Role, error) {
	r := new(Role)
	err := rm.m.get(rm.m.uri("roles", id)+rm.m.q(opts), r)
	return r, err
}

// Update a role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/patch_roles_by_id
func (rm *RoleManager) Update(id string, r *Role) (err error) {
	return rm.m.patch(rm.m.uri("roles", id), r)
}

// Delete a role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/delete_roles_by_id
func (rm *RoleManager) Delete(id string) (err error) {
	return rm.m.delete(rm.m.uri("roles", id))
}

// Retrieve a list of roles that can be assigned to users or groups.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/get_roles
func (rm *RoleManager) List(opts ...reqOption) ([]*Role, error) {
	var r []*Role
	err := rm.m.get(rm.m.uri("roles")+rm.m.q(opts), &r)
	return r, err
}

// Assign users to a role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/post_role_users
func (rm *RoleManager) AssignUsers(id string, users ...*User) error {
	u := make(map[string][]*string)
	u["users"] = make([]*string, len(users))
	for i, user := range users {
		u["users"][i] = user.ID
	}
	return rm.m.post(rm.m.uri("roles", id, "users"), &u)
}

// Retrieve users associated with a role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/get_role_user
func (rm *RoleManager) Users(id string, opts ...reqOption) ([]*User, error) {
	var u []*User
	err := rm.m.get(rm.m.uri("roles", id, "users")+rm.m.q(opts), &u)
	return u, err
}

// Associate permissions with a role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/post_role_permission_assignment
func (rm *RoleManager) AssociatePermissions(id string, permissions ...*Permission) error {
	p := make(map[string][]*Permission)
	p["permissions"] = permissions
	return rm.m.post(rm.m.uri("roles", id, "permissions"), &p)
}

// Retrieve list of permissions granted by a role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/get_role_permission
func (rm *RoleManager) Permissions(id string, opts ...reqOption) ([]*Permission, error) {
	var p []*Permission
	err := rm.m.get(rm.m.uri("roles", id, "permissions")+rm.m.q(opts), &p)
	return p, err
}

// Remove permissions associated with a role.
//
// See: https://auth0.com/docs/api/management/v2#!/Roles/delete_role_permission_assignment
func (rm *RoleManager) RemovePermissions(id string, permissions ...*Permission) error {
	p := make(map[string][]*Permission)
	p["permissions"] = permissions
	return rm.m.request("DELETE", rm.m.uri("roles", id, "permissions"), &p)
}
