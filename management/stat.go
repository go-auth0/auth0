package management

import "time"

type StatManager struct {
	m *Management
}

func NewStatManager(m *Management) *StatManager {
	return &StatManager{m}
}

// Retrieve the number of active users that logged in during the last 30 days.
//
// See: https://auth0.com/docs/api/management/v2#!/Stats/get_active_users
func (sm *StatManager) ActiveUsers() (int, error) {
	var i int
	err := sm.m.get(sm.m.uri("stats/active-users"), &i)
	return i, err
}

type DailyStat struct {
	Date            *time.Time `json:"date"`
	Logins          *int       `json:"logins"`
	Signups         *int       `json:"signups"`
	LeakedPasswords *int       `json:"leaked_passwords"`
	UpdatedAt       *time.Time `json:"updated_at"`
	CreatedAt       *time.Time `json:"created_at"`
}

func (d *DailyStat) String() string {
	return Stringify(d)
}

// Retrieve the number of logins, signups and breached-password detections
// (subscription required) that occurred each day within a specified date range.
//
// See: https://auth0.com/docs/api/management/v2#!/Stats/get_daily
func (sm *StatManager) Daily(opts ...reqOption) ([]*DailyStat, error) {
	var ds []*DailyStat
	err := sm.m.get(sm.m.uri("stats/daily")+sm.m.q(opts), &ds)
	return ds, err
}
