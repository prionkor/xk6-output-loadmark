package loadmark

import (
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"go.k6.io/k6/v2/metrics"
)

func TestWebSocketTransport(t *testing.T) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	received := make(chan []byte, 1)

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Errorf("failed to upgrade connection: %v", err)
			return
		}
		defer conn.Close()

		_, message, err := conn.ReadMessage()
		if err != nil {
			t.Errorf("failed to read message: %v", err)
			return
		}

		received <- message
	})

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to create listener: %v", err)
	}
	defer listener.Close()

	server := &http.Server{
		Handler: mux,
	}

	go func() {
		_ = server.Serve(listener)
	}()
	defer server.Close()

	cfg := Config{
		Protocol: "ws",
		URL:      "ws://" + listener.Addr().String() + "/ws",
	}

	transport, err := NewWebSocketTransport(cfg)
	if err != nil {
		t.Fatalf("failed to create transport: %v", err)
	}

	if err := transport.Start(); err != nil {
		t.Fatalf("Start() failed: %v", err)
	}

	samples := []metrics.SampleContainer{}

	if err := transport.Send(samples); err != nil {
		t.Fatalf("Send() failed: %v", err)
	}

	select {
	case message := <-received:
		if string(message) != "[]" {
			t.Fatalf("unexpected message: %s", message)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for websocket message")
	}

	if err := transport.Stop(); err != nil {
		t.Fatalf("Stop() failed: %v", err)
	}
}
