package management

import "gopkg.in/auth0.v2"

type RuleConfig struct {

	// The key for a RuleConfigs config
	Key *string `json:"key,omitempty"`

	// The value for the rules config
	Value *string `json:"value,omitempty"`
}

func (r *RuleConfig) String() string {
	return Stringify(r)
}

type RuleConfigManager struct {
	m *Management
}

func NewRuleConfigManager(m *Management) *RuleConfigManager {
	return &RuleConfigManager{m}
}

// Sets a rules config variable.
//
// See: https://auth0.com/docs/api/management/v2#!/Rules_Configs/put_rules_configs_by_key
func (rm *RuleConfigManager) Upsert(key string, r *RuleConfig) (err error) {
	return rm.m.put(rm.m.uri("rules-configs", key), r)
}

// Retrieve rules config variable keys.
//
// Note: For security, config variable values cannot be retrieved outside rule
// execution.
//
// See: https://auth0.com/docs/api/management/v2#!/Rules_Configs/get_rules_configs
func (rm *RuleConfigManager) Read(key string) (*RuleConfig, error) {
	var rs []*RuleConfig
	err := rm.m.get(rm.m.uri("rules-configs"), &rs)
	if err != nil {
		return nil, err
	}
	for _, r := range rs {
		rkey := auth0.StringValue(r.Key)
		if rkey == key {
			return r, nil
		}
	}
	return nil, &managementError{404, "Not Found", "Rule config not found"}
}

// Delete a rules config variable identified by its key.
//
// See: https://auth0.com/docs/api/management/v2#!/Rules_Configs/delete_rules_configs_by_key
func (rm *RuleConfigManager) Delete(key string) (err error) {
	return rm.m.delete(rm.m.uri("rules-configs", key))
}
