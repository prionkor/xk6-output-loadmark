## How it works

The extension receives metric samples from k6 and forwards them as JSON to a configured HTTP or WebSocket endpoint.

The destination is configured using environment variables, so the same k6 test can send results to different systems without changing the test script.

## Configuration

The extension uses the following environment variables:

```bash
XK6_OUTPUT_LOADMARK_PROTOCOL=http
XK6_OUTPUT_LOADMARK_URL=https://example.com/k6
```

### `XK6_OUTPUT_LOADMARK_PROTOCOL`

Selects the output protocol.

Supported values:

* `http`
* `ws`

### `XK6_OUTPUT_LOADMARK_URL`

Specifies the destination endpoint.

Examples:

```bash
# HTTP
XK6_OUTPUT_LOADMARK_PROTOCOL=http
XK6_OUTPUT_LOADMARK_URL=http://localhost:3000/metrics
```

```bash
# HTTPS
XK6_OUTPUT_LOADMARK_PROTOCOL=http
XK6_OUTPUT_LOADMARK_URL=https://example.com/metrics
```

```bash
# WebSocket
XK6_OUTPUT_LOADMARK_PROTOCOL=ws
XK6_OUTPUT_LOADMARK_URL=ws://localhost:3000/metrics
```

```bash
# Secure WebSocket
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

## Data format

The extension currently serializes the `metrics.SampleContainer` values received from k6 directly to JSON.

No custom metric schema or aggregation is applied by the extension.

This allows the receiving application to decide how the data should be processed, stored, or transformed.

## Error handling

The extension treats the configured output as part of the k6 test.

Configuration errors, connection failures, and output failures are reported rather than silently ignored.

For example, running the extension without a protocol results in:

```text
XK6_OUTPUT_LOADMARK_PROTOCOL is required
```

An unsupported protocol results in:

```text
unsupported protocol: ftp
```

## Development

Build the extension locally:

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
