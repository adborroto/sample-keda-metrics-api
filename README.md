# KEDA Sample for Metrics API


A simple Docker container written in **go** that will expose metrics. This metrics will be used with for a Keda scaler [Metrics APi](https://keda.sh/docs/2.3/scalers/metrics-api/)

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

# Setup

1. Deploy [Keda](https://keda.sh/docs/2.3/deploy/) in a cluster.
2. Deploy metric api 
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: metrics-deployment
  labels:
    app: metrics
spec:
  replicas: 1
  selector:
    matchLabels:
      app: metrics
  template:
    metadata:
      labels:
        app: metrics
    spec:
      containers:
      - name: metrics
        image: adoborroto/sample-keda-metrics-api:latest
        ports:
        - containerPort: 8090
        imagePullPolicy: IfNotPresent
```
3. Deploy the app that you want to scale based on the metric
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: testapp-deployment
spec:
  selector:
    matchLabels:
      app: testapp
  replicas: 1
  template:
    metadata:
      labels:
        app: testapp
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```
4. Setup the Keda metric scaler
```
apiVersion: keda.sh/v1alpha1
kind: ScaledObject
metadata:
  name: http-scaledobject
  namespace: default
  labels:
    deploymentName: http-scaled
spec:
  maxReplicaCount: 5
  scaleTargetRef:
    name: testapp-deployment
  triggers:
    - type: metrics-api
      metadata:
        targetValue: "2"
        url: http://your-service-metrics-here
        valueLocation: 'payload.value'
```