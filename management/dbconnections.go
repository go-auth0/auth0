package management

import "encoding/json"

type DBConnectionsChangePassword struct {

	// The client_id of your client. We strongly recommend including a Client ID so that the email template knows from which client the request was triggered.
	ClientID *string `json:"clientID,omitempty"`

	// The user's email address.
	Email *string `json:"email`

	// The new password. See the next paragraph for the case when a password can be set.
	Password *string `json:"password,omitempty"`

	// The name of the database connection configured to your client.
	Connection *string `json:"connection"`
}

func (c *DBConnectionsChangePassword) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return string(b)
}

type DBConnectionsManager struct {
	m *Management
}

func NewDBConnectionsManager(m *Management) *DBConnectionsManager {
	return &DBConnectionsManager{m}
}

func (dbcm *DBConnectionsManager) ChangePassword(dbChangePw *DBConnectionsChangePassword) (string, error) {

	resp, err := dbcm.m.plainHTTPRequest("POST", dbcm.m.plainURI("dbconnections", "change_password"), dbChangePw)

	return string(resp), err
}
