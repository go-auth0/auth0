package management

import (
	"fmt"
	"testing"
	"time"
	"encoding/json"
	"strings"
)

func TestConnection(t *testing.T) {

	c := &Connection{
		Name:     fmt.Sprintf("Test-Connection-%d", time.Now().Unix()),
		Strategy: "auth0",
	}

	var err error

	t.Run("Validate Contract", func(t *testing.T) {
		data := `
			{
				"id": "",
				"name": "",
				"strategy": "",
				"options": {
					"validation": {},
					"passwordPolicy": "",
					"password_history": {},
					"password_no_personal_info": {},
					"password_dictionary": {},
					"api_enable_users": true,
					"basic_profile": true,
					"ext_admin": true,
					"ext_is_suspended": true,
					"ext_agreed_terms": true,
					"ext_groups": true,
					"ext_assigned_plans": true,
					"ext_profile": true,
					"enabledDatabaseCustomization": true,
					"brute_force_protection": true,
					"import_mode": true,
					"disable_signup": true,
					"upstream_params": {},
					"client_id": "",
					"client_secret": "",
					"tenant_domain": "",
					"domain_aliases": [],
					"use_wsfed": true,
					"waad_protocol": "",
					"waad_common_endpoint": true,
					"app_id": "",
					"app_domain": "",
					"custom_scripts": {},
					"configuration": {}
				},
				"enabled_clients": [],
				"realms": [],
				"metadata": ""
			}
		`

		reader := strings.NewReader(data)
		decoder := json.NewDecoder(reader)
		decoder.DisallowUnknownFields()

		var connection Connection
		if err := decoder.Decode(&connection); err != nil {
			t.Fatal("Connection should unmarshall valid json:\n", err)
		}
	})

	t.Run("Create", func(t *testing.T) {
		err = m.Connection.Create(c)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("Read", func(t *testing.T) {
		c, err = m.Connection.Read(c.ID)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("Update", func(t *testing.T) {
		id := c.ID
		c.ID = ""       // read-only
		c.Name = ""     // read-only
		c.Strategy = "" // read-only

		err = m.Connection.Update(id, c)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("Update custom scripts and configuration values", func(t *testing.T) {
		id := c.ID
		c.ID = ""       // read-only
		c.Name = ""     // read-only
		c.Strategy = "" // read-only
		c.Options.CustomScripts = map[string]interface{}{
			"get_user": "function(email, callback) { return callback(null) }",
		}
		c.Options.Configuration = map[string]interface{}{
			"KEY_1": "VALUE_1",
			"KEY_2": "VALUE_2",
		}

		err = m.Connection.Update(id, c)
		if err != nil {
			t.Error(err)
		}

		if c.Options.CustomScripts["get_user"] != "function(email, callback) { return callback(null) }" {
			t.Fatal("response should contain a custom script with the key 'get_user'")
		}
		if value, exist := c.Options.Configuration["KEY_1"]; exist {
			if value == "VALUE_1" {
				t.Fatal("response should contain a CRYPTED configuration value for 'KEY_1'")
			}
		} else {
			t.Fatal("response should contain a configuration key 'KEY_1'")
		}
		if value, exist := c.Options.Configuration["KEY_2"]; exist {
			if value == "VALUE_1" {
				t.Fatal("response should contain a CRYPTED configuration value for 'KEY_2'")
			}
		} else {
			t.Fatal("response should contain a configuration key 'KEY_2'")
		}

	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Connection.Delete(c.ID)
		if err != nil {
			t.Error(err)
		}
	})
}
