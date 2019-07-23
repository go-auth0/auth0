package management

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnitRole(t *testing.T) {

	var err error

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields for Role", func(t *testing.T) {

		roleInJson := `{
			"id": "generated id",
			"name": "name",
			"description": "description"
		}`

		decoder := json.NewDecoder(strings.NewReader(roleInJson))
		decoder.DisallowUnknownFields()
		var role Role
		err = decoder.Decode(&role)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields for Role", func(t *testing.T) {
		role := Role{}
		bytes, _ := json.Marshal(role)
		roleInJson := string(bytes)
		wanted := `{}`
		if roleInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", roleInJson, wanted)
		}
	})

	// Based on Auth0 examples, we can verify that the definition of the struct is OK
	// by unit testing a parse of a json
	t.Run("should parse all fields for Permission", func(t *testing.T) {

		permissionInJson := `{
			"resource_server_identifier": "resource server id",
			"resource_server_name": "resource server name",
			"permission_name": "permission name",
			"description": "description"
		}`

		decoder := json.NewDecoder(strings.NewReader(permissionInJson))
		decoder.DisallowUnknownFields()
		var permission Permission
		err = decoder.Decode(&permission)
		if err != nil {
			t.Fatalf("should parse all fields. %v", err)
		}
	})

	// Based on this article https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/
	// we can verify that the use of the omitempty tag is OK
	// by unit testing a marshalling of a minimal struct
	t.Run("should omit empty fields for Role", func(t *testing.T) {
		permission := Permission{}
		bytes, _ := json.Marshal(permission)
		permissionInJson := string(bytes)
		wanted := `{}`
		if permissionInJson != wanted {
			t.Fatalf("should omit empty fields. %q, want %q", permissionInJson, wanted)
		}
	})
}
