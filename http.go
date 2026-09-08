package loadmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"go.k6.io/k6/v2/metrics"
)

type HTTPTransport struct {
	config Config
	client *http.Client
}

func NewHTTPTransport(cfg Config) (Transport, error) {
	return &HTTPTransport{
		config: cfg,
		client: &http.Client{},
	}, nil
}

func (t *HTTPTransport) Start() error {
	return nil
}

func (t *HTTPTransport) Send(samples []metrics.SampleContainer) error {
	data, err := json.Marshal(samples)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		t.config.URL,
		bytes.NewReader(data),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := t.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("server returned status %s", resp.Status)
	}

	return nil
}

func (t *HTTPTransport) Stop() error {
	return nil
}
