package management

import (
	"testing"
	"time"

	"gopkg.in/auth0.v4"
	"gopkg.in/auth0.v4/internal/testing/expect"
)

func TestLogStream(t *testing.T) {
	ls := &LogStream{
		Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
		Type: auth0.String(LogStreamSinkDatadog),
	}
	ls.Sink = &DatadogSink{
		DatadogAPIKey: auth0.String("12334567876543"),
		DatadogRegion: auth0.String("eu"),
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.LogStream.Create(ls)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := ls.Sink.(*DatadogSink); !ok {
			t.Errorf("unexpected options type %T", ls.Sink)
		}
		t.Logf("%v\n", ls)
	})

	t.Run("Read", func(t *testing.T) {
		ls, err = m.LogStream.Read(ls.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", ls)
	})

	t.Run("List", func(t *testing.T) {
		lsl, err := m.LogStream.List()
		if err != nil {
			t.Error(err)
		}
		for _, ls := range lsl {
			var ok bool

			switch ls.GetType() {
			case LogStreamSinkEventBridge:
				_, ok = ls.Sink.(*EventBridgeSink)
			case LogStreamSinkEventGrid:
				_, ok = ls.Sink.(*EventGridSink)
			case LogStreamSinkHTTP:
				_, ok = ls.Sink.(*HTTPSink)
			case LogStreamSinkDatadog:
				_, ok = ls.Sink.(*DatadogSink)
			case LogStreamSinkSplunk:
				_, ok = ls.Sink.(*SplunkSink)
			default:
				_, ok = ls.Sink.(map[string]interface{})
			}

			if !ok {
				t.Errorf("unexpected options type %T", ls.Sink)
			}

			t.Logf("%s %s %T\n", ls.GetID(), ls.GetName(), ls.Sink)
		}
	})

	t.Run("Update", func(t *testing.T) {

		id := ls.GetID()

		ls.ID = nil   // read-only
		ls.Name = nil // read-only
		ls.Type = nil // read-only

		ls.Sink = &DatadogSink{
			DatadogAPIKey: auth0.String("12334567876543"),
			DatadogRegion: auth0.String("us"),
		}

		err = m.LogStream.Update(id, ls)
		if err != nil {
			t.Error(err)
		}

		t.Logf("%v\n", ls)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.LogStream.Delete(ls.GetID())
		if err != nil {
			t.Error(err)
		}
	})
}
func TestLogStreamSinks(t *testing.T) {

	t.Run(LogStreamSinkEventBridge, func(t *testing.T) {
		g := &LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamSinkEventBridge),
			Sink: &EventBridgeSink{
				AWSAccountID: auth0.String("999999999999"),
				AWSRegion:    auth0.String("us-west-2"),
			},
		}

		defer func() { m.LogStream.Delete(g.GetID()) }()

		err := m.LogStream.Create(g)
		if err != nil {
			t.Fatal(err)
		}

		o, ok := g.Sink.(*EventBridgeSink)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}
		expect.Expect(t, g.GetStatus(), "active")
		expect.Expect(t, g.GetType(), LogStreamSinkEventBridge)
		expect.Expect(t, o.GetAWSAccountID(), "999999999999")
		expect.Expect(t, o.GetAWSRegion(), "us-west-2")
		expect.Expect(t, len(o.GetAWSPartnerEventSource()) > 0, true)

		t.Logf("%s\n", g)
	})
	/*
		t.Run(LogStreamSinkEventGrid, func(t *testing.T) {
			g := &LogStream{
				Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
				Type: auth0.String(LogStreamSinkEventGrid),
				Sink: &EventGridSink{
					AzureSubscriptionID: auth0.String("b69a6835-57c7-4d53-b0d5-1c6ae580b6d5"),
					AzureRegion:         auth0.String("northeurope"),
					AzureResourceGroup:  auth0.String("azure-logs-rg"),
				},
			}

			defer func() { m.LogStream.Delete(g.GetID()) }()

			err := m.LogStream.Create(g)
			if err != nil {
				t.Fatal(err)
			}

			o, ok := g.Sink.(*EventGridSink)
			if !ok {
				t.Fatalf("unexpected type %T", o)
			}

			expect.Expect(t, o.GetAzureSubscriptionID(), "b69a6835-57c7-4d53-b0d5-1c6ae580b6d5")
			expect.Expect(t, o.GetAzureRegion(), "northeurope")
			expect.Expect(t, o.GetAzureResourceGroup(), "azure-logs-rg")
			expect.Expect(t, len(o.GetAzurePartnerTopic()) > 0, true)

			t.Logf("%s\n", g)
		})
	*/
	t.Run(LogStreamSinkHTTP, func(t *testing.T) {
		g := &LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamSinkHTTP),
			Sink: &HTTPSink{
				HTTPEndpoint:      auth0.String("https://example.com/logs"),
				HTTPAuthorization: auth0.String("Bearer sfjkdshfkjsdhfæadkjhhags"),
				HTTPContentFormat: auth0.String("JSONLINES"),
				HTTPContentType:   auth0.String("application/json"),
			},
		}

		defer func() { m.LogStream.Delete(g.GetID()) }()

		err := m.LogStream.Create(g)
		if err != nil {
			t.Fatal(err)
		}

		o, ok := g.Sink.(*HTTPSink)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}
		expect.Expect(t, g.GetStatus(), "active")
		expect.Expect(t, g.GetType(), LogStreamSinkHTTP)
		expect.Expect(t, o.GetHTTPEndpoint(), "https://example.com/logs")
		expect.Expect(t, o.GetHTTPAuthorization(), "Bearer sfjkdshfkjsdhfæadkjhhags")
		expect.Expect(t, o.GetHTTPContentFormat(), "JSONLINES")
		expect.Expect(t, o.GetHTTPContentType(), "application/json")

		t.Logf("%s\n", g)
	})
	t.Run(LogStreamSinkDatadog, func(t *testing.T) {
		g := &LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamSinkDatadog),
			Sink: &DatadogSink{
				DatadogAPIKey: auth0.String("121233123455"),
				DatadogRegion: auth0.String("us"),
			},
		}

		defer func() { m.LogStream.Delete(g.GetID()) }()

		err := m.LogStream.Create(g)
		if err != nil {
			t.Fatal(err)
		}

		o, ok := g.Sink.(*DatadogSink)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}
		expect.Expect(t, g.GetStatus(), "active")
		expect.Expect(t, g.GetType(), LogStreamSinkDatadog)
		expect.Expect(t, o.GetDatadogAPIKey(), "121233123455")
		expect.Expect(t, o.GetDatadogRegion(), "us")

		t.Logf("%s\n", g)
	})
	t.Run(LogStreamSinkSplunk, func(t *testing.T) {
		g := &LogStream{
			Name: auth0.Stringf("Test-LogStream-%d", time.Now().Unix()),
			Type: auth0.String(LogStreamSinkSplunk),
			Sink: &SplunkSink{
				SplunkDomain: auth0.String("demo.splunk.com"),
				SplunkPort:   auth0.String("8080"),
				SplunkSecure: auth0.Bool(true),
				SplunkToken:  auth0.String("12a34ab5-c6d7-8901-23ef-456b7c89d0c1"),
			},
		}

		defer func() { m.LogStream.Delete(g.GetID()) }()

		err := m.LogStream.Create(g)
		if err != nil {
			t.Fatal(err)
		}

		o, ok := g.Sink.(*SplunkSink)
		if !ok {
			t.Fatalf("unexpected type %T", o)
		}
		expect.Expect(t, g.GetStatus(), "active")
		expect.Expect(t, g.GetType(), LogStreamSinkSplunk)
		expect.Expect(t, o.GetSplunkDomain(), "demo.splunk.com")
		expect.Expect(t, o.GetSplunkPort(), "8080")
		expect.Expect(t, o.GetSplunkSecure(), true)
		expect.Expect(t, o.GetSplunkToken(), "12a34ab5-c6d7-8901-23ef-456b7c89d0c1")
		t.Logf("%s\n", g)
	})
}
