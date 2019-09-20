# PrometheusClient

Query Prometheus from the command line.

## Usage

`PrometheusClient` will send a query to prometheus and return the response from prometheus.

Send single queries
```bash
$ prometheusClient query 'up{job="prometheus"}'
```

Send multiple queries at once
```bash
$ prometheusClient query 'avg_over_time(up{job="prometheus"})' 'up{job="prometheus"}'
```