# A dummy consumer

A toy counterpart of another toy,
[dl-kahara/dl-specimen-dummy-producer](https://github.com/dl-kahara/dl-specimen-dummy-producer).

This HTTP server responds to POST requests on `SERVER_ADDRESS`, expecting a JSON
payload, and returns a "raw" (binary) SHA256 of the received payload, to the client.

Following Prometheus-formatted counters (in addition to Golang's builtin metrics) are exposed on `METRICS_ADDRESS`:

```
datalounges_consumer_requests_total
datalounges_consumer_bytes_total
```

## Runtime configuration

```console
# Configuration takes place over environment variables, and the defaults are as follows:
SERVER_ADDRESS=:8080 \
    METRICS_ADDRESS=:9108 \
    go run .
```
