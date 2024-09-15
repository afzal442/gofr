package main

import (
	"context"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	n "gofr.dev/pkg/gofr/datasource/pubsub/nats"
	"gofr.dev/pkg/gofr/logging"
	"gofr.dev/pkg/gofr/testutil"
)

type mockMetrics struct{}

func (m *mockMetrics) IncrementCounter(ctx context.Context, name string, labels ...string) {}

// Wrapper struct for *nats.Conn that implements n.ConnInterface
type connWrapper struct {
	*nats.Conn
}

// Implement the NatsConn method for the wrapper
func (w *connWrapper) NatsConn() *nats.Conn {
	return w.Conn
}

func runNATSServer() (*server.Server, error) {
	opts := &server.Options{
		ConfigFile: "configs/nats-server.conf",
		JetStream:  true,
		Port:       -1,
		Trace:      true,
	}
	return server.NewServer(opts)
}

func TestExampleSubscriber(t *testing.T) {
	// Start the embedded NATS server
	natsServer, err := runNATSServer()
	if err != nil {
		t.Fatalf("Failed to start NATS server: %v", err)
	}
	defer natsServer.Shutdown()

	natsServer.Start()

	if !natsServer.ReadyForConnections(5 * time.Second) {
		t.Fatal("NATS server failed to start")
	}

	serverURL := natsServer.ClientURL()

	logs := testutil.StdoutOutputForFunc(func() {
		// Start the main application
		go main()

		// Wait for the application to initialize
		time.Sleep(2 * time.Second)

		// Initialize test data
		initializeTest(t, serverURL)

		// Wait for messages to be processed
		time.Sleep(5 * time.Second)
	})

	testCases := []struct {
		desc        string
		expectedLog string
	}{
		{
			desc:        "valid order",
			expectedLog: "Received order",
		},
		{
			desc:        "valid product",
			expectedLog: "Received product",
		},
	}

	for i, tc := range testCases {
		if !strings.Contains(logs, tc.expectedLog) {
			t.Errorf("TEST[%d] Failed.\n%s\nExpected log: %s\nActual logs: %s",
				i, tc.desc, tc.expectedLog, logs)
		}
	}
}

func initializeTest(t *testing.T, serverURL string, opts ...nats.Option) {
	conf := &n.Config{
		Server: serverURL,
		Stream: n.StreamConfig{
			Stream:  "sample-stream",
			Subject: "order-logs,products",
		},
	}

	mockMetrics := &mockMetrics{}
	logger := logging.NewMockLogger(logging.DEBUG)

	client, err := n.NewNATSClient(conf, logger, mockMetrics,
		func(serverURL string, opts ...nats.Option) (n.ConnInterface, error) {
			conn, err := nats.Connect(serverURL, opts...)
			if err != nil {
				return nil, err
			}
			return &connWrapper{conn}, nil
		},
		func(nc *nats.Conn) (jetstream.JetStream, error) {
			return jetstream.New(nc)
		},
	)

	if err != nil {
		t.Fatalf("Error initializing NATS client: %v", err)
	}

	ctx := context.Background()

	// Create or update stream
	streamConfig := jetstream.StreamConfig{
		Name: "sample-stream",
		// Subjects: []string{"order-logs", "products"},
	}

	log.Printf("Creating stream %s", streamConfig.Name)

	_, err = client.CreateOrUpdateStream(ctx, streamConfig)
	if err != nil {
		t.Fatalf("Error creating stream: %v", err)
	}

	// Publish test messages
	err = client.Publish(ctx, "order-logs", []byte(`{"orderId":"123","status":"pending"}`))
	if err != nil {
		t.Errorf("Error publishing to 'order-logs': %v", err)
	}

	err = client.Publish(ctx, "products", []byte(`{"productId":"123","price":"599"}`))
	if err != nil {
		t.Errorf("Error publishing to 'products': %v", err)
	}
}
