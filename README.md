# promclient

Query Prometheus from the command line.

## Usage

`promclient` will send a query to prometheus and return the response from prometheus.

Send single queries
```bash
$ promclient query 'up{job="prometheus"}'
```

Send multiple queries at once
```bash
$ promclient query 'avg_over_time(up{job="prometheus"})' 'up{job="prometheus"}'
```