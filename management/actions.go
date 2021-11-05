package management

import (
	"net/http"
	"time"
)

const (
	ActionTriggerPostLogin         string = "post-login"
	ActionTriggerClientCredentials string = "client-credentials"
)

type ActionTrigger struct {
	ID      *string `json:"id"`
	Version *string `json:"version"`
	Status  *string `json:"status,omitempty"`
}

type ActionTriggerList struct {
	Triggers []*ActionTrigger `json:"triggers"`
}

type ActionDependency struct {
	Name        *string `json:"name"`
	Version     *string `json:"version,omitempty"`
	RegistryURL *string `json:"registry_url,omitempty"`
}

type ActionSecret struct {
	Name      *string    `json:"name"`
	Value     *string    `json:"value,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ActionVersionError struct {
	ID      *string `json:"id"`
	Message *string `json:"msg"`
	Url     *string `json:"url"`
}

const (
	ActionStatusPending  string = "pending"
	ActionStatusBuilding string = "building"
	ActionStatusPackaged string = "packaged"
	ActionStatusBuilt    string = "built"
	ActionStatusRetrying string = "retrying"
	ActionStatusFailed   string = "failed"
)

type Action struct {
	// ID of the action
	ID *string `json:"id,omitempty"`
	// The name of an action
	Name *string `json:"name"`
	// List of triggers that this action supports. At this time, an action can
	// only target a single trigger at a time.
	SupportedTriggers []*ActionTrigger `json:"supported_triggers"`
	// The source code of the action.
	Code *string `json:"code,omitempty"`
	// List of third party npm modules, and their versions, that this action
	// depends on.
	Dependencies []*ActionDependency `json:"dependencies,omitempty"`
	// The Node runtime. For example `node16`, defaults to `node12`
	Runtime *string `json:"runtime,omitempty"`
	// List of secrets that are included in an action or a version of an action.
	Secrets []*ActionSecret `json:"secrets,omitempty"`
	// Version of the action that is currently deployed.
	DeployedVersion *ActionVersion `json:"deployed_version,omitempty"`
	// The build status of this action.
	Status *string `json:"status,omitempty"`
	// True if all of an Action's contents have been deployed.
	AllChangesDeployed bool `json:"all_changes_deployed,omitempty"`
	// The time when this action was built successfully.
	BuiltAt *time.Time `json:"built_at,omitempty"`
	// The time when this action was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The time when this action was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ActionList struct {
	List
	Actions []*Action `json:"actions"`
}

type ActionVersion struct {
	ID           *string             `json:"id,omitempty"`
	Code         *string             `json:"code"`
	Dependencies []*ActionDependency `json:"dependencies,omitempty"`
	Deployed     bool                `json:"deployed"`
	Status       *string             `json:"status,omitempty"`
	Number       int                 `json:"number,omitempty"`

	Errors []*ActionVersionError `json:"errors,omitempty"`
	Action *Action               `json:"action,omitempty"`

	BuiltAt   *time.Time `json:"built_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ActionVersionList struct {
	List
	Versions []*ActionVersion `json:"versions"`
}

const (
	ActionBindingReferenceByName string = "action_name"
	ActionBindingReferenceById   string = "action_id"
)

type ActionBindingReference struct {
	Type  *string `json:"type"`
	Value *string `json:"value"`
}

type ActionBinding struct {
	ID          *string `json:"id,omitempty"`
	TriggerID   *string `json:"trigger_id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`

	Ref     *ActionBindingReference `json:"ref,omitempty"`
	Action  *Action                 `json:"action,omitempty"`
	Secrets []*ActionSecret         `json:"secrets,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ActionBindingList struct {
	List
	Bindings []*ActionBinding `json:"bindings"`
}

type actionBindingsPerTrigger struct {
	Bindings []*ActionBinding `json:"bindings"`
}

type ActionTestPayload map[string]interface{}

type actionTestRequest struct {
	Payload *ActionTestPayload `json:"payload"`
}

type ActionExecutionResult struct {
	ActionName *string                `json:"action_name,omitempty"`
	Error      map[string]interface{} `json:"error,omitempty"`

	StartedAt *time.Time `json:"started_at,omitempty"`
	EndedAt   *time.Time `json:"ended_at,omitempty"`
}

type ActionExecution struct {
	ID        *string                  `json:"id"`
	TriggerID *string                  `json:"trigger_id"`
	Status    *string                  `json:"status"`
	Results   []*ActionExecutionResult `json:"results"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ActionManager struct {
	*Management
}

func newActionManager(m *Management) *ActionManager {
	return &ActionManager{m}
}

func applyActionsListDefaults(options []RequestOption) RequestOption {
	return newRequestOption(func(r *http.Request) {
		PerPage(50).apply(r)
		for _, option := range options {
			option.apply(r)
		}
	})
}

// Triggers lists the available triggers.
//
// https://auth0.com/docs/api/management/v2/#!/Actions/get_triggers
func (m *ActionManager) Triggers(opts ...RequestOption) (l *ActionTriggerList, err error) {
	err = m.Request("GET", m.URI("actions", "triggers"), &l, opts...)
	return
}

// ListTriggers lists the available triggers.
//
// Deprecated: use Triggers() instead
func (m *ActionManager) ListTriggers(opts ...RequestOption) (l *ActionTriggerList, err error) {
	return m.Triggers(opts...)
}

// Create a new action.
//
// See: https://auth0.com/docs/api/management/v2#!/Actions/post_action
func (m *ActionManager) Create(a *Action, opts ...RequestOption) error {
	return m.Request("POST", m.URI("actions", "actions"), a, opts...)
}

// Retrieve action details.
//
// See: https://auth0.com/docs/api/management/v2#!/Actions/get_action
func (m *ActionManager) Read(id string, opts ...RequestOption) (a *Action, err error) {
	err = m.Request("GET", m.URI("actions", "actions", id), &a, opts...)
	return
}

// Update an existing action.
//
// See: https://auth0.com/docs/api/management/v2#!/Actions/patch_action
func (m *ActionManager) Update(id string, a *Action, opts ...RequestOption) error {
	return m.Request("PATCH", m.URI("actions", "actions", id), &a, opts...)
}

// Delete an action
//
// See: https://auth0.com/docs/api/management/v2#!/Actions/delete_action
func (m *ActionManager) Delete(id string, opts ...RequestOption) error {
	return m.Request("DELETE", m.URI("actions", "actions", id), nil, opts...)
}

// List all actions.
//
// See: https://auth0.com/docs/api/management/v2#!/Actions/get_actions
func (m *ActionManager) List(opts ...RequestOption) (l *ActionList, err error) {
	err = m.Request("GET", m.URI("actions", "actions"), &l, applyActionsListDefaults(opts))
	return
}

// Version retrieves the version of an action.
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/get_action_version
func (m *ActionManager) Version(id string, versionId string, opts ...RequestOption) (v *ActionVersion, err error) {
	err = m.Request("GET", m.URI("actions", "actions", id, "versions", versionId), &v, opts...)
	return
}

// ReadVersion retrieves the version of an action.
//
// Deprecated: use Version() instead.
func (m *ActionManager) ReadVersion(id string, versionId string, opts ...RequestOption) (v *ActionVersion, err error) {
	return m.Version(id, versionId, opts...)
}

// Versions lists all versions of an action.
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/get_action_versions
func (m *ActionManager) Versions(id string, opts ...RequestOption) (c *ActionVersionList, err error) {
	err = m.Request("GET", m.URI("actions", "actions", id, "versions"), &c, applyActionsListDefaults(opts))
	return
}

// ListVersions of an action.
//
// Deprecated: use Versions() instead.
func (m *ActionManager) ListVersions(id string, opts ...RequestOption) (c *ActionVersionList, err error) {
	return m.Versions(id, opts...)
}

// UpdateBindings of a trigger.
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/patch_bindings
func (m *ActionManager) UpdateBindings(triggerID string, b []*ActionBinding, opts ...RequestOption) error {
	bl := &actionBindingsPerTrigger{
		Bindings: b,
	}
	return m.Request("PATCH", m.URI("actions", "triggers", triggerID, "bindings"), &bl, opts...)
}

// Bindings lists the bindings of a trigger.
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/get_bindings
func (m *ActionManager) Bindings(triggerID string, opts ...RequestOption) (bl *ActionBindingList, err error) {
	err = m.Request("GET", m.URI("actions", "triggers", triggerID, "bindings"), &bl, applyActionsListDefaults(opts))
	return
}

// ListBindings lists the bindings of a trigger.
//
// Deprecated: use Bindings() instead.
func (m *ActionManager) ListBindings(triggerID string, opts ...RequestOption) (bl *ActionBindingList, err error) {
	return m.Bindings(triggerID, opts...)
}

// Deploy an action
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/post_deploy_action
func (m *ActionManager) Deploy(id string, opts ...RequestOption) (v *ActionVersion, err error) {
	err = m.Request("POST", m.URI("actions", "actions", id, "deploy"), &v, opts...)
	return
}

// DeployVersion of an action
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/post_deploy_draft_version
func (m *ActionManager) DeployVersion(id string, versionId string, opts ...RequestOption) (v *ActionVersion, err error) {
	err = m.Request("POST", m.URI("actions", "actions", id, "versions", versionId, "deploy"), &v, opts...)
	return
}

// Test an action
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/post_test_action
func (m *ActionManager) Test(id string, payload *ActionTestPayload, opts ...RequestOption) (err error) {
	r := &actionTestRequest{
		Payload: payload,
	}
	err = m.Request("POST", m.URI("actions", "actions", id, "test"), &r, opts...)
	return
}

// Execution retrieves the details of an action execution
//
// See: https://auth0.com/docs/api/management/v2/#!/Actions/get_execution
func (m *ActionManager) Execution(executionId string, opts ...RequestOption) (v *ActionExecution, err error) {
	err = m.Request("GET", m.URI("actions", "executions", executionId), &v, opts...)
	return
}

// ReadExecution retrieves the details of an action execution
//
// Deprecated: use Execution() instead
func (m *ActionManager) ReadExecution(executionId string, opts ...RequestOption) (v *ActionExecution, err error) {
	return m.Execution(executionId, opts...)
}
