package management

type Rule struct {

	// The rule's identifier.
	ID *string `json:"id,omitempty"`

	// The name of the rule. Can only contain alphanumeric characters, spaces
	// and '-'. Can neither start nor end with '-' or spaces.
	Name *string `json:"name,omitempty"`

	// A script that contains the rule's code.
	Script *string `json:"script,omitempty"`

	// The rule's order in relation to other rules. A rule with a lower order
	// than another rule executes first. If no order is provided it will
	// automatically be one greater than the current maximum.
	Order *int `json:"order,omitempty"`

	// Enabled should be set to true if the rule is enabled, false otherwise.
	Enabled *bool `json:"enabled,omitempty"`
}

func (r *Rule) String() string {
	return Stringify(r)
}

type RuleManager struct {
	m *Management
}

func NewRuleManager(m *Management) *RuleManager {
	return &RuleManager{m}
}

// Create a new rule.
//
// Note: Changing a rule's stage of execution from the default `login_success`
// can change the rule's function signature to have user omitted.
//
// See: https://auth0.com/docs/api/management/v2#!/Rules/post_rules
func (rm *RuleManager) Create(r *Rule) error {
	return rm.m.post(rm.m.uri("rules"), r)
}

// Retrieve rule details. Accepts a list of fields to include or exclude in the result.
//
// See: https://auth0.com/docs/api/management/v2#!/Rules/get_rules_by_id
func (rm *RuleManager) Read(id string, opts ...ReqOption) (*Rule, error) {
	r := new(Rule)
	err := rm.m.get(rm.m.uri("rules", id)+rm.m.q(opts), r)
	return r, err
}

// Update an existing rule.
//
// See: https://auth0.com/docs/api/management/v2#!/Rules/patch_rules_by_id
func (rm *RuleManager) Update(id string, r *Rule) (err error) {
	return rm.m.patch(rm.m.uri("rules", id), r)
}

// Delete a rule.
//
// See: https://auth0.com/docs/api/management/v2#!/Rules/delete_rules_by_id
func (rm *RuleManager) Delete(id string) (err error) {
	return rm.m.delete(rm.m.uri("rules", id))
}
