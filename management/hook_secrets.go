package management

type HookSecrets = map[string]string

type HookSecretsManager struct {
	*Management
}

func newHookSecretsManager(m *Management) *HookSecretsManager {
	return &HookSecretsManager{m}
}

// Upserts hook secrets
//
// See: https://auth0.com/docs/api/management/v2#!/Hooks/post_secrets
func (m *HookSecretsManager) Upsert(hookId string, r *HookSecrets) (err error) {
	return m.post(m.hookPath(hookId), r)
}

// Reads hook secrets
//
// Note: For security, hook secret values cannot be retrieved outside rule
// execution (they all appear as "_VALUE_NOT_SHOWN_")
//
// See: https://auth0.com/docs/api/management/v2#!/Rules_Configs/get_rules_configs
func (m *HookSecretsManager) Read(hookId string) (r *HookSecrets, err error) {
	err = m.get(m.hookPath(hookId), &r)
	return
}

// Delete a list of hook secret keys from a given hook's secrets identified by its hookId and the keys
//
// See: https://auth0.com/docs/api/management/v2#!/Rules_Configs/delete_rules_configs_by_key
func (m *HookSecretsManager) Delete(hookId string, keys ...string) (err error) {
	return m.request("DELETE", m.hookPath(hookId), keys)
}

func (m *HookSecretsManager) hookPath(hookId string) string {
	return m.uri("hooks", hookId, "secrets")
}
