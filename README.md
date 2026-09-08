# xk6-output-loadmark

A [k6](https://k6.io/) output extension that forwards k6 metric samples to an HTTP or WebSocket endpoint.

LoadMark can use this extension to receive k6 metrics during a test, but the extension is generic and can send metrics to any compatible HTTP or WebSocket endpoint.

## Features

- HTTP and HTTPS output
- WebSocket and secure WebSocket output
- JSON output
- Configured through environment variables
- Works with standard k6 output configuration
- No custom metric schema or aggregation

## Installation

### Download a release

Pre-built k6 binaries with the LoadMark output extension are available from the project's GitHub releases.

Download the binary for your operating system and architecture, then run:

```bash
./k6 run test.js --out loadmark
```

### Build with xk6

You can also build a custom k6 binary using [xk6](https://github.com/grafana/xk6):

```bash
xk6 build --with github.com/prionkor/xk6-output-loadmark@v0.1.0
```

Or build directly from a local checkout:

```bash
xk6 build --with github.com/prionkor/xk6-output-loadmark=.
```

## Usage

The extension is enabled with the standard k6 output option:

```bash
./k6 run test.js --out loadmark
```

The destination is configured using environment variables.

## Configuration

The extension requires the following environment variables:

```bash
XK6_OUTPUT_LOADMARK_PROTOCOL=http
XK6_OUTPUT_LOADMARK_URL=https://example.com/k6
```

### `XK6_OUTPUT_LOADMARK_PROTOCOL`

Selects the output protocol.

Supported values:

- `http`
- `ws`

### `XK6_OUTPUT_LOADMARK_URL`

Specifies the destination endpoint.

HTTP:

```bash
XK6_OUTPUT_LOADMARK_PROTOCOL=http
XK6_OUTPUT_LOADMARK_URL=http://localhost:3000/metrics
```

HTTPS:

```bash
XK6_OUTPUT_LOADMARK_PROTOCOL=http
XK6_OUTPUT_LOADMARK_URL=https://example.com/metrics
```

WebSocket:

```bash
XK6_OUTPUT_LOADMARK_PROTOCOL=ws
XK6_OUTPUT_LOADMARK_URL=ws://localhost:3000/metrics
```

Secure WebSocket:

```bash
XK6_OUTPUT_LOADMARK_PROTOCOL=ws
XK6_OUTPUT_LOADMARK_URL=wss://example.com/metrics
```

Both variables are required when the `loadmark` output is used.

## HTTP output

When `http` is selected, the extension sends metric samples using HTTP `POST` requests with:

```http
Content-Type: application/json
```

The request body contains the JSON representation of the k6 metric samples received by the extension.

Example:

```bash
XK6_OUTPUT_LOADMARK_PROTOCOL=http \
XK6_OUTPUT_LOADMARK_URL=https://example.com/k6 \
./k6 run test.js --out loadmark
```

Any `2xx` HTTP response is considered successful. A non-`2xx` response is treated as an output error.

## WebSocket output

When `ws` is selected, the extension establishes a WebSocket connection to the configured URL and sends metric samples as JSON messages.

Example:

```bash
XK6_OUTPUT_LOADMARK_PROTOCOL=ws \
XK6_OUTPUT_LOADMARK_URL=wss://example.com/k6 \
./k6 run test.js --out loadmark
```

The WebSocket connection remains open while the k6 test is running and is closed when the test finishes.

## Data format

The extension currently serializes the `metrics.SampleContainer` values received from k6 directly to JSON.

No custom metric schema or aggregation is applied by the extension.

This allows the receiving application to decide how the data should be processed, stored, or transformed.

## Error handling

The extension reports configuration and output errors instead of silently ignoring them.

Missing configuration results in an error:

```text
XK6_OUTPUT_LOADMARK_PROTOCOL is required
```

An unsupported protocol results in an error:

```text
unsupported protocol: ftp
```

Connection and transport errors are also reported.

## Development

Clone the repository and install the required dependencies:

```bash
go mod download
```

Build a local k6 binary with the extension:

```bash
xk6 build --with github.com/prionkor/xk6-output-loadmark=.
```

Run tests:

```bash
go test ./...
```

Format the code:

```bash
gofmt -w .
```

## License

Apache License 2.0.
