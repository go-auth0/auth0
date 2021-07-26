package management

import (
	"fmt"
	"gopkg.in/auth0.v5"
	"os"
	"testing"
	"time"
)

func TestOrganization(t *testing.T) {
	var err error

	ts := time.Now().Format("20060102150405")

	client := &Client{
		Name:              auth0.String(fmt.Sprintf("testclient%v", ts)),
		AppType:           auth0.String("regular_web"),
		GrantTypes:        []interface{}{"client_credentials"},
		OrganizationUsage: auth0.String("allow"),
	}
	err = m.Client.Create(client)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		m.Client.Delete(auth0.StringValue(client.ClientID))
	})

	conn := &Connection{
		Name:        auth0.String(fmt.Sprintf("testconn%v", ts)),
		DisplayName: auth0.String(fmt.Sprintf("Test Connection %v", ts)),
		Strategy:    auth0.String(ConnectionStrategyAuth0),
		EnabledClients: []interface{}{
			os.Getenv("AUTH0_CLIENT_ID"),
			client.ClientID,
		},
	}
	err = m.Connection.Create(conn)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		m.Connection.Delete(conn.GetID())
	})

	user := &User{
		Connection: conn.Name,
		Email:      auth0.String("chuck@chucknorris.com"),
		Password:   auth0.String("Passwords hide their Chuck"),
		GivenName:  auth0.String("Chuck"),
		FamilyName: auth0.String("Norris"),
		Nickname:   auth0.String("Chucky"),
		UserMetadata: map[string]interface{}{
			"favourite_attack": "roundhouse_kick",
		},
		EmailVerified: auth0.Bool(true),
		VerifyEmail:   auth0.Bool(false),
		AppMetadata: map[string]interface{}{
			"facts": []string{
				"count_to_infinity_twice",
				"kill_two_stones_with_one_bird",
				"can_hear_sign_language",
				"knows_victorias_secret",
			},
		},
		Picture: auth0.String("https://example-picture-url.jpg"),
		Blocked: auth0.Bool(false),
	}
	err = m.User.Create(user)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		m.User.Delete(user.GetID())
	})

	role := &Role{
		Name:        auth0.String("admin"),
		Description: auth0.String("Administrator"),
	}
	err = m.Role.Create(role)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		m.Role.Delete(role.GetID())
	})

	o := &Organization{
		Name:        auth0.String(fmt.Sprintf("testorganization%v", ts)),
		DisplayName: auth0.String("Test Organization"),
		Branding: &OrganizationBranding{
			LogoUrl: auth0.String("https://example.com/logo.gif"),
		},
	}

	oi := &OrganizationInvitation{
		Inviter: &OrganizationInvitationInviter{
			Name: auth0.String("Test Inviter"),
		},
		Invitee: &OrganizationInvitationInvitee{
			Email: auth0.String("test@example.com"),
		},
		ClientID: client.ClientID,
	}

	oc := &OrganizationConnection{
		ConnectionID:            conn.ID,
		AssignMembershipOnLogin: auth0.Bool(true),
	}

	t.Run("Create", func(t *testing.T) {
		err = m.Organization.Create(o)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", o)
		oi.OrganizationID = o.ID
	})

	t.Run("List", func(t *testing.T) {
		var ol *OrganizationList
		ol, err = m.Organization.List()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", ol)
	})

	t.Run("ReadByID", func(t *testing.T) {
		o, err = m.Organization.Read(o.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", o)
	})

	t.Run("Update", func(t *testing.T) {
		err = m.Organization.Update(o)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", o)
	})

	t.Run("ReadByName", func(t *testing.T) {
		o, err = m.Organization.ReadByName(o.GetName())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", o)
	})

	t.Run("Connections", func(t *testing.T) {
		var l *OrganizationConnectionList
		l, err = m.Organization.Connections(o.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", l)
	})

	t.Run("AddConnection", func(t *testing.T) {
		err = m.Organization.AddConnection(o.GetID(), oc)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", oc)
	})

	t.Run("Connection", func(t *testing.T) {
		oc, err = m.Organization.Connection(o.GetID(), oc.GetConnectionID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", oc)
	})

	t.Run("UpdateConnection", func(t *testing.T) {
		err = m.Organization.UpdateConnection(o.GetID(), oc)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", oc)
	})

	t.Run("Invitations", func(t *testing.T) {
		var l *OrganizationInvitationList
		l, err = m.Organization.Invitations(o.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", l)
	})

	t.Run("CreateInvitation", func(t *testing.T) {
		err = m.Organization.CreateInvitation(oi)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", oi)
	})

	t.Run("Invitation", func(t *testing.T) {
		var i *OrganizationInvitation
		i, err = m.Organization.Invitation(o.GetID(), oi.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", i)
	})

	t.Run("Members", func(t *testing.T) {
		var l *OrganizationMemberList
		l, err = m.Organization.Members(o.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", l)
	})

	t.Run("AddMembers", func(t *testing.T) {
		err = m.Organization.AddMembers(o.GetID(), []string{user.GetID()})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("MemberRoles", func(t *testing.T) {
		var l *OrganizationMemberRoleList
		l, err = m.Organization.MemberRoles(o.GetID(), user.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", l)
	})

	t.Run("AssignMemberRoles", func(t *testing.T) {
		// (id string, userID string, roles []string) (err error)
		err = m.Organization.AssignMemberRoles(o.GetID(), user.GetID(), []string{role.GetID()})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteMemberRoles", func(t *testing.T) {
		err = m.Organization.DeleteMemberRoles(o.GetID(), user.GetID(), []string{role.GetID()})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteMember", func(t *testing.T) {
		err = m.Organization.DeleteMember(o.GetID(), []string{user.GetID()})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteConnection", func(t *testing.T) {
		err = m.Organization.DeleteConnection(o.GetID(), oc.GetConnectionID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", oc)
	})

	t.Run("DeleteInvitation", func(t *testing.T) {
		err = m.Organization.DeleteInvitation(o.GetID(), oi.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", oi)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Organization.Delete(o.GetID())
		if err != nil {
			t.Error(err)
		}
	})
}
