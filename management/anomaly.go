package management

type AnomalyManager struct {
	*Management
}

func newAnomalyManager(m *Management) *AnomalyManager {
	return &AnomalyManager{m}
}

// Check if a given IP address is blocked via the multiple user accounts
// trigger due to multiple failed logins.
//
// See: https://auth0.com/docs/api/management/v2#!/Anomaly/get_ips_by_id
func (m *AnomalyManager) CheckIP(ip string, opts ...RequestOption) (err error) {
	err = m.Request("GET",  m.URI("anomaly", "blocks", "ips", ip), nil, opts...)
	return
}

// Unblock an IP address currently blocked by the multiple user accounts
// trigger due to multiple failed logins.
//
// See: https://auth0.com/docs/api/management/v2#!/Anomaly/delete_ips_by_id
func (m *AnomalyManager) UnblockIP(ip string, opts ...RequestOption) (err error) {
	err = m.Request("DELETE", m.URI("anomaly", "blocks", "ips", ip), nil, opts...)
	return
}
