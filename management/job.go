package management

import (
	"encoding/json"
	"time"
)

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

	// The id of the connection.
	ConnectionID *string `json:"connection_id,omitempty"`
	// The url to download the result of the job.
	Location *string `json:"location,omitempty"`
	// The percentage of the work done so far.
	PercentageDone *int `json:"percentage_done,omitempty"`
	// Estimated amount of time remaining to finish the job.
	TimeLeftSeconds *int `json:"time_left_seconds,omitempty"`
	// The format of the file. Valid values are: "json" and "csv".
	Format *string `json:"format,omitempty"`
	// Limit the number of records.
	Limit *int `json:"limit,omitempty"`
	// A list of fields to be included in the CSV. If omitted, a set of
	// predefined fields will be exported.
	Fields []map[string]interface{} `json:"fields,omitempty"`
}

func (j *Job) String() string {
	b, _ := json.MarshalIndent(j, "", "  ")
	return string(b)
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

func (jm *JobManager) Read(id string, opts ...reqOption) (*Job, error) {
	j := new(Job)
	err := jm.m.get(jm.m.uri("jobs", id)+jm.m.q(opts), j)
	return j, err
}

func (jm *JobManager) ExportUsers(j *Job) error {
	return jm.m.post(jm.m.uri("jobs/users-exports"), j)
}
