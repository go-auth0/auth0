package management

import "time"

type Job struct {
	// The job's identifier. Useful to retrieve its status
	ID *string `json:"id,omitempty"`
	// The job's status
	Status *string `json:"status,omitempty"`
	// The type of job
	Type *string `json:"type,omitempty"`
	// The date when the job was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The user_id of the user to whom the email will be sent
	UserID *string `json:"user_id,omitempty"`
	// The id of the client, if not provided the global one will be used
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
