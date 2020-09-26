package management

import "encoding/json"

const (
	LogStreamSinkEventBridge = "eventbridge"
	LogStreamSinkEventGrid   = "eventgrid"
	LogStreamSinkHTTP        = "http"
	LogStreamSinkDatadog     = "datadog"
	LogStreamSinkSplunk      = "splunk"
)

type LogStream struct {
	// The hook's identifier.
	ID *string `json:"id,omitempty"`

	// The name of the hook. Can only contain alphanumeric characters, spaces
	// and '-'. Can neither start nor end with '-' or spaces.
	Name *string `json:"name,omitempty"`

	// The type of the log-stream. Can be one of "http", "eventbridge",
	// "eventgrid", "datadog" or "splunk".
	Type *string `json:"type,omitempty"`

	// The status of the log-stream. Can be one of "active", "paused", or "suspended".
	Status *string `json:"status,omitempty"`

	// Sink for validation.
	Sink    interface{}     `json:"-"`
	RawSink json.RawMessage `json:"sink,omitempty"`
}

type AWSSink struct {
	// AWS Account Id
	AWSAccountID *string `json:"awsAccountId,omitempty"`
	// AWS Region
	AWSRegion *string `json:"awsRegion,omitempty"`
	// AWS Partner Event Source
	AWSPartnerEventSource *string `json:"awsPartnerEventSource,omitempty"`
}
type AzureSink struct {
	// Azure Subscription Id
	AzureSubscriptionID *string `json:"azureSubscriptionId,omitempty"`
	// Azure Resource Group
	AzureResourceGroup *string `json:"azureResourceGroup,omitempty"`
	// Azure Region
	AzureRegion *string `json:"azureRegion,omitempty"`
	// Azure Partner Topic
	AzurePartnerTopic *string `json:"azurePartnerTopic,omitempty"`
}
type HTTPSink struct {
	// Http ContentFormat
	HTTPContentFormat *string `json:"httpContentFormat,omitempty"`
	// Http ContentType
	HTTPContentType *string `json:"httpContentType,omitempty"`
	// Http Endpoint
	HTTPEndpoint *string `json:"httpEndpoint,omitempty"`
	// Http Authorization
	HTTPAuthorization *string `json:"httpAuthorization,omitempty"`
}
type DatadogSink struct {
	// Datadog Region
	DatadogRegion *string `json:"datadogRegion,omitempty"`
	// Datadog Api Key
	DatadogAPIKey *string `json:"datadogApiKey,omitempty"`
}
type SplunkSink struct {
	// Splunk Domain
	SplunkDomain *string `json:"splunkDomain,omitempty"`
	// Splunk Token
	SplunkToken *string `json:"splunkToken,omitempty"`
	// Splunk Port
	SplunkPort *string `json:"splunkPort,omitempty"`
	// Splunk Secure
	SplunkSecure *bool `json:"splunkSecure,omitempty"`
}

type LogStreamManager struct {
	*Management
}

func newLogStreamManager(m *Management) *LogStreamManager {
	return &LogStreamManager{m}
}

func (ls *LogStream) MarshalJSON() ([]byte, error) {

	type logStream LogStream

	if ls.Sink != nil {
		b, err := json.Marshal(ls.Sink)
		if err != nil {
			return nil, err
		}
		ls.RawSink = b
	}

	return json.Marshal((*logStream)(ls))
}

func (ls *LogStream) UnmarshalJSON(b []byte) error {

	type logStream LogStream

	err := json.Unmarshal(b, (*logStream)(ls))
	if err != nil {
		return err
	}

	if ls.Type != nil {

		var v interface{}

		switch *ls.Type {
		case LogStreamSinkEventBridge:
			v = &AWSSink{}
		case LogStreamSinkEventGrid:
			v = &AzureSink{}
		case LogStreamSinkHTTP:
			v = &HTTPSink{}
		case LogStreamSinkDatadog:
			v = &DatadogSink{}
		case LogStreamSinkSplunk:
			v = &SplunkSink{}
		default:
			v = make(map[string]interface{})
		}

		err = json.Unmarshal(ls.RawSink, &v)
		if err != nil {
			return err
		}

		ls.Sink = v
	}

	return nil
}

// Create a Log Stream sink.
//
// The LogStream object requires different properties depending on the type
// of sin (which is specified using the type property):
//
// - `http` requires `httpEndpoint`, `httpContentType`, `httpContentFormat`, and `httpAuthorization`
// - `eventbridge` requires `awsAccountId`, and `awsRegion`
// - `eventgrid` requires `azureSubscriptionId`, `azureResourceGroup`, and `azureRegion`
// - `datadog` requires `datadogRegion`, and `datadogApiKey`
// - `splunk` requires `splunkDomain`, `splunkToken`, `splunkPort`, and `splunkSecure`
//
// See: https://auth0.com/docs/api/management/v2#!/log-streams
func (m *LogStreamManager) Create(e *LogStream) error {
	return m.post(m.uri("log-streams"), e)
}

// Retrieve log-stream detail by its id.
//
// See: https://auth0.com/docs/api/management/v2#!/Log_Streams/get_log_streams_by_id
func (m *LogStreamManager) Read(id string) (e *LogStream, err error) {
	err = m.get(m.uri("log-streams", id), &e)
	return
}

// List all connections.
//
// See: https://auth0.com/docs/api/management/v2#!/log-streams/get_log_streams
func (m *LogStreamManager) List() (ls []*LogStream, err error) {
	err = m.get(m.uri("log-streams"), &ls)
	return
}

// Update log-stream.
//
// See: https://auth0.com/docs/api/management/v2#!/log-streams
func (m *LogStreamManager) Update(id string, e *LogStream) (err error) {
	return m.patch(m.uri("log-streams", id), e)
}

// Delete the log-stream.
//
// See: https://auth0.com/docs/api/management/v2#!/log-streams
func (m *LogStreamManager) Delete(id string) (err error) {
	return m.delete(m.uri("log-streams", id))
}
