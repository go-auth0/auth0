package management

import (
	"encoding/json"
	"strings"
	"testing"

	"gopkg.in/auth0.v1"
)

func TestTenant(t *testing.T) {

	var tn *Tenant
	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {
		tenantInJson := `{
			"change_password": {
				"enabled": true,
				"html": "<html>Change Password</html>"
			},
			"guardian_mfa_page": {
				"enabled": true,
				"html": "<html>MFA</html>"
			},
			"default_audience": "https://mycompany.auth0.com/api/v2/",
			"default_directory": "Username-Password-Authentication",
			"error_page": {
				"html": "<html>Error Page</html>",
				"show_log_link": false,
				"url": "https://mycompany.org/error"
			},
			"flags": {
				"change_pwd_flow_v1": false,
				"enable_client_connections": false,
				"enable_apis_section": true,
				"disable_impersonation": true,
				"enable_pipeline2": false,
				"enable_dynamic_client_registration": false,
				"enable_custom_domain_in_emails": false,
				"enable_sso": false,
				"allow_changing_enable_sso": true,
				"universal_login": false,
				"enable_legacy_logs_search_v2": false,
				"disable_clickjack_protection_headers": false,
				"enable_public_signup_user_exists_error": false
			},
			"friendly_name": "My Example Tenant",
			"picture_url": "https://mycompany.org/logo.png",
			"support_email": "support@example.com",
			"support_url": "https://support.example.com",
			"allowed_logout_urls": ["https://mycompany.org/logoutCallback"],
			"session_lifetime": 168,
			"sandbox_version": "8",
			"sandbox_versions_available": ["8"],
			"idle_session_lifetime": 72,
			"universal_login": {
				"colors": {
					"primary": "#0076BC",
					"page_background": "#000000"
				}
			},
			"enabled_locales": [ "en" ],
			"device_flow": {
				"charset": "base20",
				"mask": "****-****"
			}
		}`

		decoder := json.NewDecoder(strings.NewReader(tenantInJson))
		decoder.DisallowUnknownFields()
		var tenant Tenant
		err = decoder.Decode(&tenant)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		tenant := Tenant{}
		bytes, _ := json.Marshal(tenant)
		tenantInJson := string(bytes)
		wanted := "{}"
		if tenantInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", tenantInJson, wanted)
		}
	})

	t.Run("Read", func(t *testing.T) {
		tn, err = m.Tenant.Read()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", tn)
	})

	t.Run("Update", func(t *testing.T) {
		err = m.Tenant.Update(&Tenant{
			FriendlyName: auth0.String("My Example Tenant"),
			SupportURL:   auth0.String("https://support.example.com"),
			SupportEmail: auth0.String("support@example.com"),
		})
		if err != nil {
			t.Error(err)
		}
		tn, _ = m.Tenant.Read()
		t.Logf("%v\n", tn)
	})
}
