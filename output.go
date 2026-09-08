package loadmark

import (
	"errors"
	"fmt"
	"os"

	"go.k6.io/k6/v2/metrics"
	"go.k6.io/k6/v2/output"
)

func init() {
	output.RegisterExtension("loadmark", New)
}

type Output struct {
	transport Transport
	err       error
}

type Transport interface {
	Start() error
	Send([]metrics.SampleContainer) error
	Stop() error
}

func New(params output.Params) (output.Output, error) {
	cfg, err := NewConfig()
	if err != nil {
		return nil, err
	}

	transport, err := NewTransport(cfg)
	if err != nil {
		return nil, err
	}

	return &Output{
		transport: transport,
	}, nil
}

func NewTransport(cfg Config) (Transport, error) {
	if cfg.Protocol == "" {
		return nil, nil
	}

	switch cfg.Protocol {
	case "http":
		return NewHTTPTransport(cfg)
	case "ws":
		return NewWebSocketTransport(cfg)
	default:
		return nil, fmt.Errorf("unsupported protocol: %s", cfg.Protocol)
	}
}

func (o *Output) Description() string {
	return "LoadMark output"
}

func (o *Output) Start() error {
	return o.transport.Start()
}

func (o *Output) AddMetricSamples(
	samples []metrics.SampleContainer,
) {
	if len(samples) == 0 || o.err != nil {
		return
	}

	if err := o.transport.Send(samples); err != nil {
		o.err = err
	}
}

func (o *Output) Stop() error {
	if err := o.transport.Stop(); err != nil {
		if o.err == nil {
			o.err = err
		}
	}

	return o.err
}

type Config struct {
	Protocol string
	URL      string
}

func NewConfig() (Config, error) {
	protocol := os.Getenv("XK6_OUTPUT_LOADMARK_PROTOCOL")
	url := os.Getenv("XK6_OUTPUT_LOADMARK_URL")

	if protocol == "" {
		return Config{}, errors.New(
			"XK6_OUTPUT_LOADMARK_PROTOCOL is required",
		)
	}

	if url == "" {
		return Config{}, errors.New(
			"XK6_OUTPUT_LOADMARK_URL is required",
		)
	}

	return Config{
		Protocol: protocol,
		URL:      url,
	}, nil
}
