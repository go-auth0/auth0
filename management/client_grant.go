package management

type ClientGrant struct {

	// A generated string identifying the client grant.
	ID *string `json:"id,omitempty"`

	// The identifier of the client.
	ClientID *string `json:"client_id,omitempty"`

	// The audience.
	Audience *string `json:"audience,omitempty"`

	Scope []interface{} `json:"scope"`
}

func (c *ClientGrant) String() string {
	return Stringify(c)
}

type ClientGrantManager struct {
	m *Management
}

func NewClientGrantManager(m *Management) *ClientGrantManager {
	return &ClientGrantManager{m}
}

// Create a client grant.
//
// See: https://auth0.com/docs/api/management/v2#!/Client_Grants/post_client_grants
func (cg *ClientGrantManager) Create(g *ClientGrant) (err error) {
	return cg.m.post(cg.m.uri("client-grants"), g)
}

// Retrieves a client grant by its id.
//
// The Auth0 Management API does not offer a method to retrieve a client grant
// by id, we fake this by listing all client grants and matching by id on the
// client side. For this reason this method should be used with caution.
func (cg *ClientGrantManager) Read(id string) (*ClientGrant, error) {
	var gs []*ClientGrant
	err := cg.m.get(cg.m.uri("client-grants"), &gs)
	if err != nil {
		return nil, err
	}
	for _, g := range gs {
		gid := *g.ID
		if gid == id {
			return g, nil
		}
	}
	return nil, &managementError{
		StatusCode: 404,
		Err:        "Not Found",
		Message:    "Client grant not found",
	}
}

// Update a client grant.
//
// See: https://auth0.com/docs/api/management/v2#!/Client_Grants/patch_client_grants_by_id
func (cg *ClientGrantManager) Update(id string, g *ClientGrant) (err error) {
	return cg.m.patch(cg.m.uri("client-grants", id), g)
}

// Delete a client grant.
//
// See: https://auth0.com/docs/api/management/v2#!/Client_Grants/delete_client_grants_by_id
func (cg *ClientGrantManager) Delete(id string) (err error) {
	return cg.m.delete(cg.m.uri("client-grants", id))
}

// Retrieve client grants.
//
// See: https://auth0.com/docs/api/management/v2#!/Client_Grants/get_client_grants
func (cg *ClientGrantManager) List(opts ...ReqOption) (gs []*ClientGrant, err error) {
	err = cg.m.get(cg.m.uri("client-grants")+cg.m.q(opts), &gs)
	return
}
