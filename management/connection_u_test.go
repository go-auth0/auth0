package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitConnection(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields", func(t *testing.T) {
		connectiorInJson := `{
			"id": "generated id",
			"name": "connection name",
			"strategy": "auth0",
			"is_domain_connection": true,
			"options": {
				"validation": {
					"key": "value"
				},
				"passwordPolicy": "low",
				"password_history": {
					"enable": false,
					"size": 5
				},
				"password_no_personal_info": {
					"enable": false
				},
				"password_dictionary": {
					"enable": false,
					"dictionary": []
				},
				"api_enable_users": true,
				"basic_profile": true,
				"ext_admin": true,
				"ext_is_suspended": true,
				"ext_agreed_terms": true,
				"ext_groups": true,
				"ext_nested_groups": true,
				"ext_assigned_plans": true,
				"ext_profile": true,
				"enabledDatabaseCustomization": true,
				"brute_force_protection": true,
				"import_mode": true,
				"disable_signup": true,
				"requires_username": true,
				"upstream_params": {
					"key": "value"
				},
				"client_id": "client id",
				"client_secret": "client secret",
				"tenant_domain": "tenant domain",
				"domain_aliases": [
					"domain alias 1",
					"domain alias 2"
				],
				"use_wsfed": true,
				"waad_protocol": "waad protocol",
				"waad_common_endpoint": true,
				"app_id": "app id",
				"app_domain": "app domain",
				"max_groups_to_retrieve": "max groups to retrieve",
				"customScripts": {
					"login": "",
					"get_user": ""
				},
				"configuration": {
					"key": "value"
				},
				"totp": {
					"time_step": 123,
					"length": 12
				},
				"name": "name",
				"twilio_sid": "twilio sid",
				"twilio_token": "twilio token",
				"from": "from",
				"syntax": "syntax",
				"template": "template",
				"messaging_service_sid": "messaging service sid",
				"adfs_server": "adfs server"
			},
			"enabled_clients": [
				"client_id_1",
				"client_id_2"
			],
			"realms": [
				"connection name"
			],
			"metadata": {
				"key": "value"
			}
		}`

		decoder := json.NewDecoder(strings.NewReader(connectiorInJson))
		decoder.DisallowUnknownFields()
		var connection Connection
		err = decoder.Decode(&connection)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields", func(t *testing.T) {
		connection := Connection{}
		bytes, _ := json.Marshal(connection)
		connectionInJson := string(bytes)
		wanted := "{}"
		if connectionInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", connectionInJson, wanted)
		}
	})
}
