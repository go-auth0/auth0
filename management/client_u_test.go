package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitClient(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {
		clientInJson := `{
			"name": "client name",
			"description": "description",
			"client_id": "client_id",
			"client_secret": "client_secret",
			"app_type": "spa",
			"logo_uri": "https://mycompany.org/logo.png",
			"is_first_party": true,
			"is_token_endpoint_ip_header_trusted": true,
			"oidc_conformant": true,
			"callbacks": ["https://mycompany.org/callback"],
			"allowed_origins": ["https://mycompany.org"],
			"web_origins": ["https://mycompany.org"],
			"client_aliases": ["https://mycompany.org"],
			"allowed_clients": ["https://mycompany.org"],
			"allowed_logout_urls": ["https://mycompany.org/logout"],
			"jwt_configuration": {
				"lifetime_in_seconds": 123,
				"secret_encoded": true,
				"scopes": [
					"scope_1",
					"scope_2"
				],
				"alg": "RS256"
			},
			"signing_keys": [
				{
					"cert": "-----BEGIN CERTIFICATE-----\r\n-----END CERTIFICATE-----\r\n",
					"pkcs7": "-----BEGIN PKCS7-----\r\n-----END PKCS7-----\r\n",
					"subject": "/CN="
				}
			],
			"encryption_key": {
				"pub": "pub",
				"cert": "cert",
				"subject": "subject"
			},
			"sso": true,
			"sso_disabled": true,
			"cross_origin_auth": true,
			"grant_types": [
				"authorization_code",
				"implicit",
				"refresh_token"
			],
			"cross_origin_loc": "https://mycompany.org/",
			"custom_login_page_on": true,
			"custom_login_page": "html page",
			"custom_login_page_preview": "html page",
			"form_template": "form template",
			"addons": {
				"samlp": {
					"lifetimeInSeconds": 3600,
					"signatureAlgorithm": "rsa-sha256",
					"digestAlgorithm": "sha256"
				}
			},
			"token_endpoint_auth_method": "none",
			"client_metadata": {
				"key": "value"
			},
			"mobile": {
				"android": {
					"app_package_name": "com.example",
					"sha256_cert_fingerprints": [
						"D8:A0:83:..."
					]
				},
				"ios": {
					"team_id": "9JA89QQLNQ",
					"app_bundle_identifier": "com.my.bundle.id"
				}
			}
		}`

		decoder := json.NewDecoder(strings.NewReader(clientInJson))
		decoder.DisallowUnknownFields()
		var client Client
		err = decoder.Decode(&client)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		client := Client{}
		bytes, _ := json.Marshal(client)
		clientInJson := string(bytes)
		wanted := "{}"
		if clientInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", clientInJson, wanted)
		}
	})
}
