package loadmark

import "go.k6.io/k6/v2/metrics"

type WebSocketTransport struct {
	config Config
}

func NewWebSocketTransport(cfg Config) (Transport, error) {
	return &WebSocketTransport{
		config: cfg,
	}, nil
}

func (t *WebSocketTransport) Start() error {
	return nil
}

func (t *WebSocketTransport) Send(samples []metrics.SampleContainer) error {
	return nil
}

func (t *WebSocketTransport) Stop() error {
	return nil
}
