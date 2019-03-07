package management

import "time"

type Job struct {
	ID        *string    `json:"id,omitempty"`
	Status    *string    `json:"status,omitempty"`
	Type      *string    `json:"type,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`

	UserID   *string `json:"user_id,omitempty"`
	ClientID *string `json:"cliend_id,omitempty"`
}

type JobManager struct {
	m *Management
}

func NewJobManager(m *Management) *JobManager {
	return &JobManager{m}
}

func (jm *JobManager) VerifyEmail(j *Job) error {
	return jm.m.post(jm.m.uri("jobs/verification-email"), j)
}
