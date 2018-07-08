package management

import "testing"

func TestTenant(t *testing.T) {

	var tn *Tenant
	var err error

	t.Run("Read", func(t *testing.T) {
		tn, err = m.Tenant.Read()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%#v\n", tn)
	})

	t.Run("Update", func(t *testing.T) {
		err = m.Tenant.Update(&Tenant{
			FriendlyName: "My Example Tenant",
			SupportURL:   "https://support.example.com",
			SupportEmail: "support@example.com",
		})
		if err != nil {
			t.Error(err)
		}
		tn, _ = m.Tenant.Read()
		t.Logf("%v\n", tn)
	})
}
