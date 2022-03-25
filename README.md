# KEDA Sample for Metrics API


A simple Docker container written in go that will expose metrics.

```
$ curl http:localhost:8090/api/value
```
Sample response:
```
{
    payload: {
        value: 0
    },
    success: true,
    error: ""
}
```

And one endpoint to update the metric.

```
$ curl --location --request POST 'http://localhost:8090/api/value/10'
```

## Docker image

You can access the image by pulling from DockerHub.

```
docker pull adoborroto/sample-keda-metrics-api:latest
```

# Test Keda