package management

import "time"

type VerificationEmailJob struct {
	UserID   string `json:"user_id"`
	ClientID string `json:"cliend_id,omitempty"`
}

type VerificationEmailJobResponse struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type JobManager struct {
	m *Management
}

func NewJobManager(m *Management) *JobManager {
	return &JobManager{m}
}

func (jm *JobManager) SendVerificationEmail(j VerificationEmailJob) (VerificationEmailJobResponse, error) {
	res := VerificationEmailJobResponse{}
	err := jm.m.request("POST", jm.m.uri("jobs/verification-email"), j, &res)
	return res, err
}
