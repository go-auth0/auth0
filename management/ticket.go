package management

type Ticket struct {
	// The user will be redirected to this endpoint once the ticket is used
	ResultURL *string `json:"result_url,omitempty"`

	// The UserID for which the ticket is to be created
	UserID *string `json:"user_id,omitempty"`

	// The ticket's lifetime in seconds starting from the moment of creation.
	// After expiration the ticket can not be used to verify the users's email.
	// If not specified or if you send 0 the Auth0 default lifetime will be
	// applied.
	TTLSec *int `json:"ttl_sec,omitempty"`

	// ID of the client. If provided for tenants using New Universal Login experience,
	// the user will be prompted to redirect to the default login route of the
	// corresponding application once the ticket is used.
	//
	// Conflicts with: ResultURL
	ClientID *string `json:"client_id,omitempty"`

	// The connection that provides the identity for which the password is to be
	// changed. If sending this parameter, the email is also required and the
	// UserID is invalid.
	//
	// Requires: Email
	// Conflicts with: UserID
	ConnectionID *string `json:"connection_id,omitempty"`

	// The user's email
	//
	// Requires: ConnectionID
	// Conflicts with: UserID
	Email *string `json:"email,omitempty"`

	// The URL that represents the ticket
	Ticket *string `json:"ticket,omitempty"`

	// Whether to set the email_verified attribute to true (true) or whether it
	// should not be updated
	MarkEmailAsVerified *bool `json:"mark_email_as_verified,omitempty"`

	// Whether to include the email address as part of the returnUrl in the reset_email (true), or not (false - default).
	IncludeEmailInRedirect *bool `json:"includeEmailInRedirect,omitempty"`
}

type TicketManager struct {
	*Management
}

func newTicketManager(m *Management) *TicketManager {
	return &TicketManager{m}
}

// VerifyEmail creates a ticket to verify a user's email address.
//
// See: https://auth0.com/docs/api/management/v2#!/Tickets/post_email_verification
func (m *TicketManager) VerifyEmail(t *Ticket, opts ...RequestOption) error {
	return m.Request("POST", m.URI("tickets", "email-verification"), t, opts...)
}

// ChangePassword creates a password change ticket for a user.
//
// See: https://auth0.com/docs/api/management/v2#!/Tickets/post_password_change
func (m *TicketManager) ChangePassword(t *Ticket, opts ...RequestOption) error {
	return m.Request("POST", m.URI("tickets", "password-change"), t, opts...)
}
