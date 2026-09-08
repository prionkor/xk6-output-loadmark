package loadmark

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"go.k6.io/k6/v2/metrics"
)

type WebSocketTransport struct {
	config Config
	conn   *websocket.Conn
}

func NewWebSocketTransport(cfg Config) (Transport, error) {
	return &WebSocketTransport{
		config: cfg,
	}, nil
}

func (t *WebSocketTransport) Start() error {
	conn, _, err := websocket.DefaultDialer.Dial(t.config.URL, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to websocket: %w", err)
	}

	t.conn = conn

	return nil
}

func (t *WebSocketTransport) Send(
	samples []metrics.SampleContainer,
) error {
	data, err := json.Marshal(samples)
	if err != nil {
		return err
	}

	if err := t.conn.WriteMessage(
		websocket.TextMessage,
		data,
	); err != nil {
		return fmt.Errorf("failed to send websocket message: %w", err)
	}

	return nil
}

func (t *WebSocketTransport) Stop() error {
	if t.conn == nil {
		return nil
	}

	return t.conn.Close()
}
